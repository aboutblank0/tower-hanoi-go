package hanoi

const TOWER_HEIGHT = 7
const DISC_MAX_WIDTH = TOWER_HEIGHT

type Tower struct {
	Height int
	Discs  []*Disc
}

type Disc struct {
	Width int
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
