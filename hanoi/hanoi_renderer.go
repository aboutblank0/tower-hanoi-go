package hanoi

import (
	"fmt"
	"strings"
)

const TOWER_PADDING = 2
func RenderHanoiGame(game HanoiGame) {
	padding := strings.Repeat(" ", TOWER_PADDING)

	for y := range TOWER_HEIGHT {
		line := new(strings.Builder)
		for _, tower := range game.Towers {
			line.WriteString(getTowerTextAtY(*tower, y))
			line.WriteString(padding)
		}

		fmt.Println(line.String())
	}
}

func getTowerTextAtY(tower Tower, y int) string {
	if y >= len(tower.Discs) {
		return ""
	}

	disc := tower.Discs[y]
	return getDiscText(disc)
}

func getDiscText(disc Disc) string {
	width := disc.Width * 2 -1
	maxWidth := DISC_MAX_WIDTH * 2 - 1

	padding := strings.Repeat(" ", (maxWidth - width)/2)
	text := padding + strings.Repeat("■", width) + padding
	return text
}
