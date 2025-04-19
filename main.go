package main

import (
	"aboutblank0/hanoi/hanoi"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func eraseScreen()    { fmt.Print("\033[2J") } //Erase Screen
func moveCursorHome() { fmt.Print("\033[H") }

func main() {
	game := hanoi.NewGame()
	hanoi.BFS(game)
	printStats()
}

func getRandInt(exclude int) int {
	for {
		r := rand.Intn(3)
		if r != exclude {
			return r
		}
	}
}

func randomGameLoop(g *hanoi.HanoiGame) {
	game := *g
	eraseScreen()
	for {
		rSrc := getRandInt(-1)
		rDest := getRandInt(rSrc)

		if !game.Move(rSrc, rDest) {
			continue
		}
		moveCursorHome()
		hanoi.RenderHanoiGame(game)

		if game.IsComplete() {
			fmt.Printf("You won win %v moves.\n", game.MovesMade)
			break
		}
		time.Sleep(time.Millisecond * 5)
	}
}

func readGameLoop(game *hanoi.HanoiGame) {
	reader := bufio.NewReader(os.Stdin)
	hanoi.RenderHanoiGame(*game)

	for running := true; running; {
		fmt.Print("Enter your move: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			continue
		}

		source := int(text[0] - '0')
		dest := int(text[1] - '0')
		if game.Move(source, dest) {
			hanoi.RenderHanoiGame(*game)
			if game.IsComplete() {
				fmt.Printf("You won in %v moves\n", game.MovesMade)
				break
			}
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}

func printStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v KB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc = %v KB\n", m.TotalAlloc/1024)
	fmt.Printf("Sys = %v KB\n", m.Sys/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}
