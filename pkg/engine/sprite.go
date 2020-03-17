package engine

import (
	"math"
	"sort"

	"github.com/godstanis/goom3d/pkg/screen"
)

// Sprite alignment
const (
	CenterAlign = iota
	TopAlign
	BottomAlign
)

// Sprite represents one sprite object on a map
type Sprite struct {
	X, Y    float64
	Align   int
	Scale   float64
	Texture [][]uint32
	Solid   bool
}

// Sprites information
var Sprites []*Sprite

// Draws sprites column on screen
func drawSpritesColumn(screen screen.Screen, col int, dir Vector, distanceToWall float64) {
	// We should sort all our sprites by distance so closest are rendered last
	sort.Slice(Sprites, func(i, j int) bool {
		return cPlayerDistToSprite(*Sprites[i]) > cPlayerDistToSprite(*Sprites[j])
	})

	for _, sprite := range Sprites {
		// Draw sprite column if it is not behind any walls
		if cPlayerDistToSprite(*sprite) < distanceToWall {
			drawSpriteColumn(screen, *sprite, col, dir)
		}
	}
}

// Draws one column for specific sprite on screen
func drawSpriteColumn(screen screen.Screen, sprite Sprite, col int, dir Vector) {
	sees, sP := seesSprite(dir, sprite)
	if !sees {
		return
	}

	baseSpriteH := distToHeight(cPlayerDistToSprite(sprite), screen.Height()) / 2
	scaledSpriteH := int(float64(baseSpriteH) * sprite.Scale)
	if scaledSpriteH == 0 {
		return
	}

	spriteScreenRow := calculateSpriteStart(screen, sprite, baseSpriteH)
	spriteCol := int(sP * float64(len(sprite.Texture[0])))
	scaledTexture := scaleTextureVertically(sprite.Texture, scaledSpriteH)

	for i := 0; i <= scaledSpriteH; i++ {
		spriteScreenRow++
		if spriteScreenRow < 0 || scaledTexture[i][spriteCol] == 0 {
			continue
		}

		_ = screen.SetPixel(col, spriteScreenRow, scaledTexture[i][spriteCol])
	}
}

// Calculates where on screen a sprite starts rendering (may return negative value, meaning starting point is UP out of screen bounds)
func calculateSpriteStart(screen screen.Screen, sprite Sprite, baseSpriteHeight int) int {
	switch sprite.Align {
	case BottomAlign:
		return screen.Height()/2 + int(float64(baseSpriteHeight)*(1-sprite.Scale))
	case CenterAlign:
		return screen.Height()/2 - int(float64(baseSpriteHeight)*sprite.Scale)/2
	case TopAlign:
		return screen.Height()/2 - baseSpriteHeight
	default:
		return 0
	}
}

// Detects if a sprite on specific angle is visible to the camera
//
// spritePoint - where on sprite's plane is hit (0.0-1.0)
func seesSprite(dir Vector, sprite Sprite) (hit bool, spritePoint float64) {
	angleDiff := playerDirToSprite(dir, sprite)
	angleOffset := 18 / playerDistToSprite(sprite) * sprite.Scale

	if math.Abs(angleDiff) < angleOffset {
		return true, (angleDiff + angleOffset) / (angleOffset * 2)
	}

	return false, 0
}

// Calculates angle between the direction and the sprite (negative is to right) in a relation with player position
func playerDirToSprite(dir Vector, sprite Sprite) float64 {
	relX, relY := sprite.X-curX, sprite.Y-curY // Coordinates of sprite relative to player's coordinates
	return math.Atan2(relX*dir.Y-relY*dir.X, relX*dir.X+relY*dir.Y) * (180 / math.Pi)
}

// Calculates distance between player and the sprite
func playerDistToSprite(sprite Sprite) float64 {
	return distToSprite(curX, curY, sprite)
}

// Calculates perpendicular corrected distance between player and the sprite
func cPlayerDistToSprite(sprite Sprite) float64 {
	correction := math.Cos(playerDirToSprite(curVector, sprite) * math.Pi / 180)
	return playerDistToSprite(sprite) * correction
}

// Calculates distance to sprite from the given points
func distToSprite(x, y float64, sprite Sprite) float64 {
	return math.Sqrt((sprite.X-x)*(sprite.X-x) + (sprite.Y-y)*(sprite.Y-y))
}

// Calculates if the given point intersects with any sprite
func intersectsWithSprite(x, y float64) (intersects bool) {
	for _, sprite := range Sprites {
		if sprite.Solid && distToSprite(x, y, *sprite) < sprite.Scale/2 {
			return true
		}
	}
	return false
}
