package main

import (
	"aboutblank0/hanoi/hanoi"
)

func main() {
	game := hanoi.NewGame()

	game.Move(0, 2)
	game.Move(0, 1)
	game.Move(2, 1)
	game.Move(0, 2)
	game.Move(1, 0)
	game.Move(1, 2)
	game.Move(0, 2)
	game.Move(0, 1)
	game.Move(2, 1)
	game.Move(2, 0)
	game.Move(1, 0)
	game.Move(2, 1)
	game.Move(0, 2)
	game.Move(0, 1)
	game.Move(2, 1)
	game.Move(0, 2)

	hanoi.RenderHanoiGame(*game)
}
