package hanoi

import "math/rand"

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
	Color DiscColor
	Width int
}

type DiscColor uint8;
const (
	ColorRed = DiscColor(iota)
	ColorGreen
	ColorYellow

	EndColor //For getting a random Color
)

func NewGame() *HanoiGame {
	game := new(HanoiGame)
	game.Towers[0] = getRandomTower()
	game.Towers[1] = getRandomTower()
	game.Towers[2] = getRandomTower()

	return game
}

func getRandomTower() *Tower {
	tower := new(Tower)
	tower.Height = TOWER_HEIGHT
	tower.Discs = make([]Disc, 0, tower.Height)
	
	for range tower.Height {
		disc := new(Disc)
		disc.Width = rand.Intn(tower.Height) + 1 //Random Size (1 to max height of tower)
		disc.Color = DiscColor(rand.Intn(int(EndColor)))
		tower.Discs = append(tower.Discs, *disc)
	}

	return tower
}
