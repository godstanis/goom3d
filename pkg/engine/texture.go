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

// Scales texture column to integer h proportionally
// Returns scaled column of texture tiles
//
// It's a simple and rough but effective texture scaling method, we just calculate relation of our original size to required
// 	For example if our texture is only 5 symbols high but we want to render it on screen with 10 height we just get the relation
// 	in this case is 5/10=0.5 of our screen row per texture row (i.e. 1 texture row for each screen row)
func scaleTileTextureCol(tile, h int, tileP float64) (scaled []string) {
	c:= int(math.Round(float64(len(TileTextures[tile][0])-1) * tileP))
	textureHeight := len(TileTextures[tile])
	scaledTexture := make([]string, h+1)

	// Actual scaling
	perTile := float64(textureHeight-1) / float64(h-1)
	curTileS := 0.0
	for i := 0; i < h; i++ {
		scaledTexture[i] = TileTextures[tile][int(math.Round(curTileS))][c]
		curTileS += perTile
	}

	// Keep first and last pixel of original texture, just to make it consistent and pleasing to look at larger distances
	if len(scaledTexture) > 0 && textureHeight > 0 {
		scaledTexture[0] = TileTextures[tile][0][c]
		scaledTexture[len(scaledTexture)-1] = TileTextures[tile][textureHeight-1][c]
	}

	return scaledTexture
}
