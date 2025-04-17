package hanoi

import (
	"fmt"
	"strings"
)

const TOWER_PADDING = 2
func RenderHanoiGame(game HanoiGame) {
	padding := strings.Repeat(" ", TOWER_PADDING)

	for y := range TOWER_HEIGHT {
		line := ""
		for _, tower := range game.Towers {
			line += getTowerTextAtHeight(*tower, y) + padding
		}

		fmt.Println(line)
	}
}

func getTowerTextAtHeight(tower Tower, height int) string {
	if height > len(tower.Discs) {
		return ""
	}

	disc := tower.Discs[height]

	return getDiscText(disc)
}

func getDiscText(disc Disc) string {
	width := disc.Width * 2 -1
	maxWidth := DISC_MAX_WIDTH * 2 - 1

	padding := strings.Repeat(" ", (maxWidth - width)/2)
	text := padding + strings.Repeat("■", width) + padding

	switch disc.Color {
	case ColorRed:
		return "\033[31m" + text + "\033[0m"
	case ColorGreen:
		return "\033[32m" + text + "\033[0m"
	case ColorYellow:
		return "\033[33m" + text + "\033[0m"
	default:
		return text
	}
}
