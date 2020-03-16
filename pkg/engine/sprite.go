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
func drawSpritesColumn(screen screen.Screen, col int, angle Degree, distanceToWall float64) {
	// We should sort all our sprites by distance so closest are rendered last
	sort.Slice(Sprites, func(i, j int) bool {
		return playerDistToSprite(*Sprites[i]) > playerDistToSprite(*Sprites[j])
	})

	for _, sprite := range Sprites {
		// Draw sprite column if it is not behind any walls
		if cPlayerDistToSprite(*sprite) < distanceToWall {
			drawSpriteColumn(screen, *sprite, col, angle)
		}
	}
}

// Draws one column for specific sprite on screen
func drawSpriteColumn(screen screen.Screen, sprite Sprite, col int, angle Degree) {
	sees, sP := seesSprite(angle, sprite)
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
func seesSprite(angle Degree, sprite Sprite) (hit bool, spritePoint float64) {
	angleDiff := angle.Plus(-playerAngleToSprite(sprite).Get())
	angleOffset := 18 / playerDistToSprite(sprite) * sprite.Scale

	if angleDiff > 0 && angleDiff < angleOffset {
		return true, (angleDiff + angleOffset) / (angleOffset * 2)
	}
	if angleDiff > 360-angleOffset {
		angleDiff = angleDiff - 360
		return true, (angleDiff + angleOffset) / (angleOffset * 2)
	}
	return false, 0
}

// Calculates angle between a player and the sprite
func playerAngleToSprite(sprite Sprite) Degree {
	relX, relY := sprite.X-curX, sprite.Y-curY // Coordinates of sprite relative to player's coordinates
	// We are calculating cos between (x:1;y:0) unit vector(represents players relation to the sprite) and the sprite's angle
	// It's a simple angle-between-two-vectors minimized-for-zero-vector formula
	smallestAngle := math.Acos(relX/math.Sqrt((relX*relX)+(relY*relY))) * (180 / math.Pi)

	// Because angle between vector is the smallest one
	// we should additionally perform this correction
	if relY <= 0 {
		return Degree{}.NewDegree(360.0 - smallestAngle)
	}
	return Degree{}.NewDegree(smallestAngle)
}

// Calculates distance between player and the sprite
func playerDistToSprite(sprite Sprite) float64 {
	return distToSprite(curX, curY, sprite)
}

// Calculates perpendicular corrected distance between player and the sprite
func cPlayerDistToSprite(sprite Sprite) float64 {
	vectorToPlayerView := curVector.NewRotated(-playerAngleToSprite(sprite).Get())
	correction := math.Cos(vectorToPlayerView.Angle().Get() * math.Pi / 180)
	return playerDistToSprite(sprite) * correction
}

// Calculates distance to sprite from the given points
func distToSprite(x, y float64, sprite Sprite) float64 {
	return math.Sqrt((sprite.X-x)*(sprite.X-x) + (sprite.Y-y)*(sprite.Y-y))
}

// Calculates if the given point intersects with any sprite
func intersectsWithSprite(x, y float64) (intersects bool) {
	for _, sprite := range Sprites {
		if sprite.Solid {
			if distToSprite(x, y, *sprite) < sprite.Scale/2 {
				return true
			}
		}
	}
	return false
}
