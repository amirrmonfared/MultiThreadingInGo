package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.Add(Vector2D{viewRadius, viewRadius}), b.position.Add(Vector2D{-viewRadius, -viewRadius})
	avgPosition, avgVelocity, seperation := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
	count := 0.0

	rwLock.RLock()
	for i := math.Max(lower.x, 0); i < math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j < math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					seperation = seperation.Add(b.position.Subtract(boids[otherBoidId].position).DivisonV(dist))
				}
			}
		}
	}
	rwLock.RUnlock()

	accel := Vector2D{b.borderBounce(b.position.x, screenWidth), b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		avgPosition, avgVelocity = avgVelocity.DivisonV(count), avgVelocity.DivisonV(count)
		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
		accelCohesion := avgPosition.Subtract(b.position).MultiplyV(adjRate)
		accelSeperation := seperation.MultiplyV(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeperation)
	}

	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne() {
	accel := b.calcAcceleration()

	rwLock.Lock()
	defer rwLock.Unlock()

	b.velocity = b.velocity.Add(accel).limit(-1, 1)

	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id

}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		id:       bid,
	}
	boids[bid] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	go b.start()
}
