// Texture mapping related stuff
package engine

import (
	"math"
)

// Texture sets for tile values
var TileTextures = map[int][][]string{
	1: {
		{"▒", "▓", "▓", "▓", "▓", "▓", "▒"},
		{"░", "░", "░", "▒", "░", "░", "░"},
		{"░", "░", "░", "▒", "░", "░", "░"},
		{"░", "░", "░", "▒", "░", "░", "░"},
		{"▒", "▒", "▒", "▓", "▒", "▒", "▒"},
		{"░", "░", "░", "▒", "░", "░", "░"},
		{"░", "░", "░", "▒", "░", "░", "░"},
		{"░", "░", "░", "▒", "░", "░", "░"},
		{"▒", "▓", "▓", "▓", "▓", "▓", "▒"},
	},
	2: {
		{"╔", "═", "═", "═", "═", "═", "═", "═", "═", "╗"},
		{"║", "░", "░", "░", "▒", "▒", "░", "░", "░", "║"},
		{"║", "░", "░", "░", "▒", "▒", "░", "░", "░", "║"},
		{"║", "░", "░", "░", "▒", "▒", "░", "░", "░", "║"},
		{"║", "░", "░", "░", "▒", "▒", "░", "░", "░", "║"},
		{"║", "▒", "▒", "▒", "▓", "▓", "▒", "▒", "▒", "║"},
		{"║", "▒", "▒", "▒", "▓", "▓", "▒", "▒", "▒", "║"},
		{"║", "░", "░", "░", "▒", "▒", "░", "░", "░", "║"},
		{"║", "░", "░", "░", "▒", "▒", "░", "░", "░", "║"},
		{"║", "░", "░", "░", "▒", "▒", "░", "░", "░", "║"},
		{"║", "░", "░", "░", "▒", "▒", "░", "░", "░", "║"},
		{"╚", "═", "═", "═", "═", "═", "═", "═", "═", "╝"},
	},
	3: {
		{"┌", "─", "─", "─", "─", "╥", "─", "─", "─", "─", "┐"},
		{"│", "░", "░", "░", "░", "╬", "░", "░", "░", "░", "│"},
		{"│", "░", "░", "░", "░", "╬", "░", "░", "░", "░", "│"},
		{"│", "░", "░", "░", "░", "╬", "░", "░", "░", "░", "│"},
		{"│", "▒", "▒", "▒", "▒", "╬", "▒", "▒", "▒", "▒", "│"},
		{"│", "▓", "▓", "▓", "▓", "╬", "▓", "▓", "▓", "▓", "│"},
		{"│", "▓", "▓", "▓", "▓", "╬", "▓", "▓", "▓", "▓", "│"},
		{"│", "▒", "▒", "▒", "▒", "╬", "▒", "▒", "▒", "▒", "│"},
		{"│", "░", "░", "░", "░", "╬", "░", "░", "░", "░", "│"},
		{"│", "░", "░", "░", "░", "╬", "░", "░", "░", "░", "│"},
		{"│", "░", "░", "░", "░", "╬", "░", "░", "░", "░", "│"},
		{"└", "─", "─", "─", "─", "╨", "─", "─", "─", "─", "┘"},
	},
	4: {
		{"▒", "▒", "▒", "▒", "▒", "▒", "▒"},
		{"▒", "░", "░", "░", "░", "░", "▒"},
		{"▒", "░", " ", " ", " ", "░", "▓"},
		{"▒", "░", " ", " ", " ", "░", "▒"},
		{"▒", "░", "░", "░", "░", "░", "▒"},
		{"▒", "░", "╬", "░", "░", "░", "▒"},
		{"▒", "░", "░", "░", "░", "░", "▓"},
		{"▒", "░", "░", "░", "░", "░", "▒"},
		{"▒", "▒", "▒", "▒", "▒", "▒", "▒"},
	},
}

// Scales 2d texture vertically to fit integer h proportionally
// Returns scaled 2d texture
//
// It's a simple and rough but effective texture scaling method, we just calculate relation of our original size to required
// 	For example if our texture is only 5 symbols high but we want to render it on screen with 10 height we just get the relation
// 	in this case is 5/10=0.5 of our screen row per texture row (i.e. 1 texture row for each screen row)
func scaleStringTextureVertically(texture [][]string, h int) (scaled [][]string) {
	scaled = make([][]string, h)

	perTile := float64(len(texture)-1) / float64(h-1)
	curTileS := 0.0
	for i := 0; i < h; i++ {
		scaled[i] = texture[int(math.Round(curTileS))]
		curTileS += perTile
	}

	return scaled
}
