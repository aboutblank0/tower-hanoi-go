package hanoi

type HanoiGame struct {
	Towers [3]*Tower
}

const TOWER_HEIGHT = 7
type Tower struct {
	Height int
	Discs []Disc
}

const DISC_MAX_WIDTH = TOWER_HEIGHT
type Disc struct {
	Width int
}

func NewGame() *HanoiGame {
	game := new(HanoiGame)
	game.Towers[0] = getInitialTower()
	game.Towers[1] = getEmptyTower()
	game.Towers[2] = getEmptyTower()

	return game
}

func getInitialTower() *Tower{
	tower := new(Tower)
	tower.Height = TOWER_HEIGHT
	tower.Discs = make([]Disc, 0, tower.Height)

	for i := range tower.Height {
		disc := new(Disc)
		disc.Width = i + 1
		tower.Discs = append(tower.Discs, *disc)
	}

	return tower
}

func getEmptyTower() *Tower {
	tower := new(Tower)
	tower.Height = TOWER_HEIGHT
	tower.Discs = make([]Disc, 0, tower.Height)
	return new(Tower)
}
