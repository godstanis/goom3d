// Texture mapping related stuff
package engine

import (
	"math"
)

// Texture sets for tile values
var TileTextures = map[int][][]uint32{}

// Scales 2d texture vertically to fit integer h proportionally
// Returns scaled 2d texture
//
// It's a simple and rough but effective texture scaling method, we just calculate relation of our original size to required
// 	For example if our texture is only 5 symbols high but we want to render it on screen with 10 height we just get the relation
// 	in this case is 5/10=0.5 of our screen row per texture row (i.e. 1 texture row for each screen row)
func scaleTextureVertically(texture [][]uint32, h int) (scaled [][]uint32) {
	scaled = make([][]uint32, h+1)

	rel := float64(len(texture)-1) / float64(h)
	for i := 0; i <= h; i++ {
		scaled[i] = texture[int(float64(i)*rel)]
	}

	return scaled
}

// The same as one above but with additional max height clipping
//
// If max height is more then the requested one - the scaled texture will not take up all returning slice, leaving empty (nil) rows
func scaleAndClipTextureVertically(texture [][]uint32, h, maxH int) (scaled [][]uint32) {
	scaled = make([][]uint32, maxH)

	var offset, tOffset int
	if h < maxH {
		offset = (maxH - h) / 2
	} else {
		tOffset = (h - maxH) / 2
	}

	rel := float64(len(texture)-1) / float64(h)
	for i := offset; i < maxH-offset; i++ {
		scaled[i] = texture[int(math.Round(float64((i-offset)+tOffset)*rel))]
	}

	return scaled
}
