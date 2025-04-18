package main

import (
	"aboutblank0/hanoi/hanoi"
	"fmt"
	"runtime"
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

func printStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v KB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc = %v KB\n", m.TotalAlloc/1024)
	fmt.Printf("Sys = %v KB\n", m.Sys/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}
