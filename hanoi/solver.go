package hanoi

import (
	"aboutblank0/hanoi/collection"
	"fmt"
	"time"
)

type Move struct {
	From int
	To   int
}

//TODO: Instead of cloning/storing different game states, have just one and mutate it around
func BFS(game *HanoiGame) {
	queue := collection.NewQueue[*HanoiGame]()
	visited := make(map[uint64]bool, 1024)

	startTime := time.Now()
	counter := 1

	start := game.Clone()
	startKey := start.Serialize()
	visited[startKey] = true
	queue.Enqueue(start)
	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()
		counter++

		if current.IsComplete() {
			fmt.Printf("Complete in %v moves!\nTime taken: %s\nBranches explored: %v\n", current.MovesMade, time.Since(startTime), counter)
			return
		}

		moves := GetAllPossibleMoves(*current)
		for _, move := range moves {
			next := current.Clone()
			next.Move(move.From, move.To)

			key := next.Serialize()
			if !visited[key] {
				visited[key] = true

				queue.Enqueue(next)
			}
		}
	}
}

var moves = make([]Move, 0, 6)
func GetAllPossibleMoves(game HanoiGame) []Move {
	moves = moves[:0]
	for i, tower := range game.Towers {
		if tower.isEmpty() {
			continue
		}
		for j := range game.Towers {
			if j == i {
				continue
			}
			if game.CanMove(i, j) {
				moves = append(moves, Move{From: i, To: j})
			}
		}
	}
	return moves
}
