# Tower of Hanoi Solver (Go)

This is a simple Tower of Hanoi game written in Go, featuring an efficient **Breadth-First Search (BFS)** solver.

![alacritty_SqsGt2J17B](https://github.com/user-attachments/assets/91974540-afc1-4d69-8ddf-89af21b6a225)


## Features

- Solves the Tower of Hanoi puzzle using BFS
- Can Handle up to 32 discs (good luck), after which the solver will break due to overflow on serialization
- Solves the 16-disc puzzle in approximately **12 seconds** (on my machine)

## Purpose

The goal of this project was to implement and optimize a BFS-based solver for the Tower of Hanoi problem, focusing on performance and clarity. All of course just for fun/learning.

## Usage

To run the solver:

```bash
go run main.go
```

To change the disc amount change `TOWER_HEIGHT` in `tower.go`
You can also manually play the game, just change the main function to call `readGameLoop()` instead

