package hanoi

import (
	"errors"
)

type HanoiGame struct {
	Towers [3]*Tower
}

func NewGame() *HanoiGame {
	game := new(HanoiGame)
	game.Towers[0] = newInitialTower()
	game.Towers[1] = newEmptyTower()
	game.Towers[2] = newEmptyTower()

	return game
}

// Moving logic in tower of hanoi is very simple.
// Since the top disc always gets moved, we just need to know
// the source and destination TOWERs (their index)
func (game *HanoiGame) Move(source int, destination int) bool {
	if source == destination {
		return false
	}

	if source >= len(game.Towers) || destination >= len(game.Towers) {
		return false
	}

	srcTower := game.Towers[source]
	destTower := game.Towers[destination]
	if !canMove(*srcTower, *destTower) {
		return false
	}

	disc, success := srcTower.pop()
	if !success {
		panic(errors.New("Attempting to pop from an empty tower"))
	}
	success = destTower.push(disc)
	if !success {
		panic(errors.New("Attempting to push into full tower"))
	}

	return true
}

func canMove(src Tower, dest Tower) bool {
	if src.isEmpty() {
		return false
	}
	if dest.isEmpty() {
		return true
	}
	return dest.peek() > src.peek()
}
