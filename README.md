# Multithreading In Go

Multi-threading examples, including a boids simulation in Go Lang

This source code has examples of creating threads and inter process communication (IPC) via
memory sharing and message passing (such as channels). It also has examples of thread 
synchronization, such as mutexes, wait groups and conditional variables.

This project uses Ebiten, A dead simple 2D game library in Go for the boids examples.
https://github.com/hajimehoshi/ebiten  
Found the library is very easy to install and use. Check out some of the games bundled with the lib.  
The installation instructions for ebiten can be found here (choose your OS):
https://ebiten.org/documents/install.html


To run any of the code examples/scenarios do the following:
```
cd <goprojects>/multithreadingingo
go build ./...
go run ./<example>
```
such as:
```
go run ./boids
```

Please do get in touch if you have any suggestions/improvements!

Follow me on https://linkedin.com/in/amirrmonfared