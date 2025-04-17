package hanoi

import (
	"fmt"
	"strings"
)

func RenderHanoiGame(game HanoiGame) {
	for _, tower := range game.Towers {
		renderTower(*tower)
		fmt.Println("=== TOWER ===")
	}
}

func renderTower(tower Tower) {
	for _, disc := range tower.Discs {
		fmt.Println(getDiscText(disc))
	}
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
