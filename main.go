package main

import (
	"aboutblank0/hanoi/hanoi"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

func eraseScreen()    { fmt.Print("\033[2J") }
func moveCursorHome() { fmt.Print("\033[H") }

func main() {
	game := hanoi.NewGame()
	moves := hanoi.BFS(game)
	replayMoves(game, moves)
	// printStats()
}

func replayMoves(g *hanoi.HanoiGame, moves []hanoi.Move) {
	game := *g
	eraseScreen()
	for _, move := range moves {
		if !game.Move(move.From, move.To) {
			fmt.Println("Skipping Move")
			continue
		}
		moveCursorHome()
		hanoi.RenderHanoiGame(game)

		time.Sleep(time.Millisecond * 10)
	}
	fmt.Printf("You won in %v moves.\n", game.MovesMade)
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
