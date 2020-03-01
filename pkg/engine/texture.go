// Texture mapping related stuff
package engine

import "math"

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
	c := translateToTextureCol(tile, tileP)
	var textureCol, scaledTexture []string
	for _, val := range TileTextures[tile] {
		textureCol = append(textureCol, val[c])
	}

	// No need for scaling if required height is equal to texture's original height
	if len(textureCol)-1 == h-1 {
		return textureCol
	}

	// Actual scaling
	perTile := float64(len(textureCol)-1) / float64(h-1)
	curTileS := 0.0
	for i := 0; i < h; i++ {
		scaledTexture = append(scaledTexture, TileTextures[tile][int(math.Round(curTileS))][c])
		curTileS += perTile
	}

	// Keep first and last pixel of original texture, just to make it consistent and pleasing to look at larger distances
	if len(scaledTexture) > 0 && len(textureCol) > 0 {
		scaledTexture[0] = textureCol[0]
		scaledTexture[len(scaledTexture)-1] = textureCol[len(textureCol)-1]
	}

	return scaledTexture
}

// Translates percentages to concrete texture column
func translateToTextureCol(tile int, tileP float64) int {
	textureWidth := len(TileTextures[tile][0])
	return int(math.Round(float64(textureWidth-1) * tileP))
}
