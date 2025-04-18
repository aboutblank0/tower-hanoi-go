package hanoi

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

	disc := srcTower.pop()
	destTower.push(disc)
	return true
}

func canMove(source Tower, dest Tower) bool {
	if !dest.hasSpace() {
		return false
	}

	topDisc, success := dest.peek()
	if !success {
		return true
	}

	srcDisc, success := source.peek()
	if !success {
		return false
	}

	return topDisc.Width > srcDisc.Width
}

