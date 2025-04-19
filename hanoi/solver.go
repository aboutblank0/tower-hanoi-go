package hanoi

import (
	"aboutblank0/hanoi/collection"
	"fmt"
)

func BFS(game *HanoiGame) {
	queue := collection.NewQueue[*HanoiGame]()
	visited := map[string]bool{}

	start := game.Clone()
	startKey := start.Serialize()
	visited[startKey] = true
	queue.Enqueue(start)

	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()

		if current.IsComplete() {
			fmt.Printf("Complete in %v moves!\n", current.MovesMade)
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

type Move struct {
	From int
	To   int
}
func GetAllPossibleMoves(game HanoiGame) []Move {
	moves := make([]Move, 0, 3)
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
