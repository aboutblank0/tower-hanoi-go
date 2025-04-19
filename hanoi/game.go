package hanoi

import (
	"errors"
)

type HanoiGame struct {
	Towers    [3]*Tower
	MovesMade int
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
func (game *HanoiGame) Move(src int, dest int) bool {
	if !game.CanMove(src, dest) {
		return false
	}

	srcTower := game.Towers[src]
	destTower := game.Towers[dest]

	disc, success := srcTower.pop()
	if !success {
		panic(errors.New("Attempting to pop from an empty tower"))
	}
	success = destTower.push(disc)
	if !success {
		panic(errors.New("Attempting to push into full tower"))
	}

	game.MovesMade++
	return true
}

func (game *HanoiGame) CanMove(src int, dest int) bool {
	if src == dest {
		return false
	}
	if src >= len(game.Towers) || dest >= len(game.Towers) {
		return false
	}

	srcTower := game.Towers[src]
	if srcTower.isEmpty() {
		return false
	}
	destTower := game.Towers[dest]
	if destTower.isEmpty() {
		return true
	}
	return destTower.peek() > srcTower.peek()
}

func (game HanoiGame) IsComplete() bool {
	return len(game.Towers[2].Discs) == TOWER_HEIGHT
}

func (game *HanoiGame) Clone() *HanoiGame {
	towers := [3]*Tower{
		game.Towers[0].Clone(),
		game.Towers[1].Clone(),
		game.Towers[2].Clone(),
	}

	return &HanoiGame{
		Towers:    towers,
		MovesMade: game.MovesMade,
	}
}

func (game *HanoiGame) Serialize() uint64 {
	var key uint64
	for towerIndex, tower := range game.Towers {
		for _, disc := range tower.Discs {
			shift := disc * 2
			key |= uint64(towerIndex) << shift
		}
	}
	return key
}
