package hanoi

import (
	"fmt"
	"strconv"
	"strings"
)

func RenderHanoiGame(game HanoiGame) {
	padding := strings.Repeat(" ", TOWER_PADDING)

	for y := TOWER_HEIGHT - 1; y >= 0; y-- {
		line := new(strings.Builder)
		for _, tower := range game.Towers {
			line.WriteString(getTowerTextAtY(*tower, y))
			line.WriteString(padding)
		}
		fmt.Println(line.String())
	}

	//Rendering the tower indexes
	line := new(strings.Builder)
	for i := range game.Towers {
		line.WriteString(getCenteredIndex(i))
		line.WriteString(padding)
	}

	fmt.Println(line.String())
}

const TOWER_PADDING = 2
const PADDING_STRING = " "
const DISC_STRING = "■"
const MAX_WIDTH = DISC_MAX_WIDTH*2 - 1

func getTowerTextAtY(tower Tower, y int) string {
	if y >= len(tower.Discs) {
		return getDiscText(0)
	}

	disc := tower.Discs[y]
	return getDiscText(disc)
}

func getDiscText(discWidth int) string {
	if discWidth == 0 {
		return strings.Repeat(PADDING_STRING, MAX_WIDTH)
	}

	width := discWidth*2 - 1
	padding := strings.Repeat(PADDING_STRING, (MAX_WIDTH-width)/2)
	text := padding + strings.Repeat(DISC_STRING, width) + padding
	return text
}

func getCenteredIndex(index int) string {
	padding := strings.Repeat(PADDING_STRING, (MAX_WIDTH-1)/2)
	text := padding + strconv.Itoa(index) + padding
	return text
}
