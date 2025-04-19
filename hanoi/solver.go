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

func (m Move) invert() Move { return Move{From: m.To, To: m.From} }

type State struct {
	Game   *HanoiGame
	Parent *State
	Move   Move
}

// TODO: Instead of cloning/storing different game states, have just one and mutate it around
func BFS(game *HanoiGame) []Move {
	queue := collection.NewQueue[*State]()
	visited := make(map[uint64]bool, 1024)

	startTime := time.Now()
	counter := 1

	startState := &State{Game: game.Clone()}
	startKey := startState.Game.Serialize()
	visited[startKey] = true
	queue.Enqueue(startState)
	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()
		counter++

		if current.Game.IsComplete() {
			fmt.Printf("Complete in %v moves!\nTime taken: %s\nBranches explored: %v\n", current.Game.MovesMade, time.Since(startTime), counter)
			return retraceMoves(*current)
		}

		moves := getAllPossibleMoves(*current.Game)
		for _, move := range moves {
			next := &State{Game: current.Game.Clone(), Parent: current, Move: move}
			next.Game.Move(move.From, move.To)

			key := next.Game.Serialize()
			if !visited[key] {
				visited[key] = true

				queue.Enqueue(next)
			}
		}
	}

	return nil
}

var moves = make([]Move, 0, 6)

func getAllPossibleMoves(game HanoiGame) []Move {
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

func retraceMoves(state State) []Move {
	moves := make([]Move, 0, 8)

	for cur := &state; cur.Parent != nil; cur = cur.Parent {
		moves = append(moves, cur.Move)
	}
	return reverse(moves)
}

func reverse(m []Move) []Move {
	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}
	return m
}
