package engine

import (
	"fmt"
	"github.com/godstanis/goom3d/pkg/screen"
)

// Draws the UI
func drawUI(screen screen.Screen) {
	drawMinimap(screen)
	drawCrosshair(screen)
}

// Draws simple crosshair
func drawCrosshair(screen screen.Screen) {
	crosshairTexture := [][]string{
		{"", "", "", "", ""},
		{"", "", "", "", ""},
		{"-", "", "", "", "-"},
		{"", "", "", "", ""},
		{"", "", "|", "", ""},
	}
	offsetX := (screen.Width() / 2) - len(crosshairTexture[0])/2
	offsetY := (screen.Height() / 2) - len(crosshairTexture)/2
	for rI, row := range crosshairTexture {
		for tI, tile := range row {
			if tile != "" {
				_ = screen.SetPixel(tI+offsetX, rI+offsetY, fmt.Sprintf("%s", tile))
			}

		}
	}
}

// Draws minimap with player position
func drawMinimap(screen screen.Screen) {
	offset := 2
	for rI, row := range Map {
		for tI, tile := range row {
			if int(curY) == rI && int(curX) == tI {
				_ = screen.SetPixel(tI+offset, rI+offset, fmt.Sprintf("%s", "▓"))
				continue
			}
			if tile == 0 {
				_ = screen.SetPixel(tI+offset, rI+offset, " ")
			} else {
				_ = screen.SetPixel(tI+offset, rI+offset, "░")
			}

		}
	}
}
