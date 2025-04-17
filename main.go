package main

import (
	"aboutblank0/hanoi/hanoi"
)

func main() {
	game := hanoi.NewGame()
	hanoi.RenderHanoiGame(*game)
}
