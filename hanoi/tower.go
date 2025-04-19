package hanoi

const TOWER_HEIGHT = 10		//MAXIMUM IS 32 For solver to work
const DISC_MAX_WIDTH = TOWER_HEIGHT

type Tower struct {
	Discs  []int
}

func (t Tower) isEmpty() bool {
	return len(t.Discs) == 0
}

// Slightly different to the usual push, since I don't want to add more than the
// original capacity. However, I still want to support multiple capacities (tower heights)
func (t *Tower) push(disc int) bool {
	if len(t.Discs) == cap(t.Discs) {
		return false
	}
	t.Discs = append(t.Discs, disc)
	return true
}

func (t *Tower) pop() (int, bool) {
	if len(t.Discs) == 0 {
		return -1, false
	}

	index := len(t.Discs) - 1
	val := t.Discs[index]
	t.Discs = t.Discs[:index]
	return val, true
}

func (t *Tower) peek() int {
	if (len(t.Discs) == 0) {
		return 0
	}
	return t.Discs[len(t.Discs)-1]
}

func newInitialTower() *Tower {
	tower := new(Tower)
	tower.Discs = make([]int, 0, TOWER_HEIGHT)

	for i := TOWER_HEIGHT - 1; i >= 0; i-- {
		tower.push(i + 1)
	}
	return tower
}

func newEmptyTower() *Tower {
	tower := new(Tower)
	tower.Discs = make([]int, 0, TOWER_HEIGHT)
	return tower
}

func (t *Tower) Clone() *Tower {
	discs := make([]int, len(t.Discs), cap(t.Discs))
	copy(discs, t.Discs)

	return &Tower{
		Discs: discs,
	}
}
