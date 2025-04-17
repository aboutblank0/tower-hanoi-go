package hanoi

type HanoiGame struct {
	Towers [3]*Tower
}

const TOWER_HEIGHT = 7

type Tower struct {
	Height int
	Discs  []*Disc
}

func (t *Tower) hasSpace() bool {
	for i := range len(t.Discs) {
		if t.Discs[i] == nil {
			return true
		}
	}
	return false
}

func (t *Tower) peek() (d Disc, success bool) {
	i := t.getTopDiscIndex()
	if i == -1 {
		return Disc{}, false
	}
	disc := t.Discs[i]
	return *disc, true
}

func (t *Tower) pop() *Disc {
	i := t.getTopDiscIndex()
	if i == -1 {
		return nil
	}
	disc := t.Discs[i]
	t.Discs[i] = nil
	return disc
}

func (t *Tower) push(disc *Disc) bool {
	for i := len(t.Discs) - 1; i >= 0; i-- {
		if t.Discs[i] == nil {
			t.Discs[i] = disc
			return true
		}
	}
	return false
}

func (t *Tower) getTopDiscIndex() int {
	for i := range t.Discs {
		disc := t.Discs[i]
		if disc != nil {
			return i
		}
	}

	return -1
}

const DISC_MAX_WIDTH = TOWER_HEIGHT

type Disc struct {
	Width int
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

func newInitialTower() *Tower {
	tower := new(Tower)
	tower.Height = TOWER_HEIGHT
	tower.Discs = make([]*Disc, tower.Height)

	for i := range tower.Height {
		disc := new(Disc)
		disc.Width = i + 1
		tower.Discs[i] = disc
	}

	return tower
}

func newEmptyTower() *Tower {
	tower := new(Tower)
	tower.Height = TOWER_HEIGHT
	tower.Discs = make([]*Disc, tower.Height)
	return tower
}
