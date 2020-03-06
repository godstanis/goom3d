package engine

import (
	"github.com/godstanis/goom3d/pkg/screen"
	"math"
	"sort"
)

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
}

// sprites information for map
var sprites = []*Sprite{
	&lampSprite,
	&boxSprite,
	&box2Sprite,
	&paintingSprite,
	&tableSprite,
	&chairSprite,
	&human,
}

var lampSprite = Sprite{
	X: 1.4, Y: 2.5, Scale: 1, Texture: [][]uint32{
		{0, 0, 0, 0x967639, 0x967639, 0, 0, 0},
		{0, 0, 0, 0x967639, 0x967639, 0, 0, 0},
		{0xE5872A, 0, 0, 0xE5DA2A, 0xE5DA2A, 0, 0, 0xE5872A},
		{0xE5DA2A, 0xE5DA2A, 0, 0xE5DA2A, 0xE5DA2A, 0, 0xE5DA2A, 0xE5DA2A},
		{0, 0, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0, 0},
		{0, 0, 0, 0xE5DA2A, 0xE5DA2A, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	},
	Align: TopAlign,
}

var boxSprite = Sprite{
	X: 1.2, Y: 1.2, Scale: 1, Texture: [][]uint32{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0, 0},
		{0, 0x967639, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0x967639, 0},
		{0x967639, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0x967639},
		{0x967639, 0xE5DA2A, 0xE5DA2A, 0x458925, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0x967639},
		{0, 0x967639, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0x967639, 0},
		{0, 0, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	},
	Align: BottomAlign,
}

var box2Sprite = Sprite{
	X: 1.5, Y: 1.7, Scale: 1, Texture: [][]uint32{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0, 0},
		{0, 0x967639, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0x967639, 0},
		{0x967639, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0x458925, 0xE5DA2A, 0x967639},
		{0x967639, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0x967639},
		{0, 0x967639, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0xE5DA2A, 0x967639, 0},
		{0, 0, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	},
	Align: BottomAlign,
}

var tableSprite = Sprite{
	X: 1.4, Y: 2.5, Scale: 1, Texture: [][]uint32{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d},
		{0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x967639},
		{0x967639, 0, 0x967639, 0, 0x967639, 0, 0x967639},
		{0x967639, 0, 0x967639, 0, 0x967639, 0, 0x967639},
		{0x967639, 0, 0x967639, 0, 0x967639, 0, 0x967639},
		{0x967639, 0, 0, 0, 0x967639, 0, 0},
	},
	Align: BottomAlign,
}

var chairSprite = Sprite{
	X: 1.4, Y: 2.1, Scale: 1, Texture: [][]uint32{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0x216583, 0},
		{0, 0, 0, 0x216583, 0},
		{0, 0, 0, 0x216583, 0},
		{0x216583, 0x216583, 0x216583, 0x216583, 0},
		{0x216583, 0, 0, 0x216583, 0},
		{0x216583, 0, 0, 0x216583, 0},
	},
	Align: BottomAlign,
}

var paintingSprite = Sprite{
	X: 2.55, Y: 1.41, Scale: 1, Texture: [][]uint32{
		{0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d},
		{0x785e2d, 0x656F84, 0x656F84, 0x525A6A, 0x525A6A, 0x525A6A, 0x65C0BA, 0x65C0BA, 0x65C0BA, 0x525A6A, 0x785e2d},
		{0x785e2d, 0x656F84, 0x656F84, 0x525A6A, 0x525A6A, 0x525A6A, 0x216583, 0x7D89A3, 0x525A6A, 0x525A6A, 0x785e2d},
		{0x785e2d, 0x656F84, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x7D89A3, 0x525A6A, 0x7D89A3, 0x785e2d},
		{0x785e2d, 0x656F84, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x7D89A3, 0x7D89A3, 0x7D89A3, 0x525A6A, 0x785e2d},
		{0x785e2d, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x7D89A3, 0x525A6A, 0x7D89A3, 0x525A6A, 0x525A6A, 0x785e2d},
		{0x785e2d, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x7D89A3, 0x525A6A, 0x525A6A, 0x785e2d},
		{0x785e2d, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x7D89A3, 0x525A6A, 0x7D89A3, 0x525A6A, 0x785e2d},
		{0x785e2d, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x7D89A3, 0x525A6A, 0x7D89A3, 0x525A6A, 0x785e2d},
		{0x785e2d, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0x525A6A, 0xF76262, 0x525A6A, 0xF76262, 0x525A6A, 0x785e2d},
		{0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d, 0x785e2d},
	},
	Align: CenterAlign,
}

var human = Sprite{
	X: 3.2, Y: 5, Scale: 2, Texture: [][]uint32{
		{0xE5872A, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xE5872A},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0x458925, 0x458925, 0x458925, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0x458925, 0xF76262, 0x458925, 0x458925, 0xF76262, 0x458925, 0, 0, 0, 0, 0},
		{0, 0, 0, 0x458925, 0x458925, 0x458925, 0x458925, 0x458925, 0x458925, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0x458925, 0x458925, 0x458925, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0x458925, 0x458925, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0x458925, 0x458925, 0x458925, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0x458925, 0x458925, 0x458925, 0x458925, 0x458925, 0x458925, 0, 0, 0, 0, 0},
		{0, 0, 0, 0x458925, 0, 0x458925, 0x458925, 0, 0x458925, 0, 0, 0, 0, 0},
		{0, 0, 0, 0x458925, 0, 0x458925, 0x458925, 0, 0x458925, 0, 0, 0, 0, 0},
		{0, 0, 0, 0x458925, 0, 0x458925, 0x458925, 0, 0x458925, 0, 0, 0, 0, 0},
		{0, 0, 0, 0x458925, 0, 0x458925, 0x458925, 0, 0x458925, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0x458925, 0x458925, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0x458925, 0x458925, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0x458925, 0x458925, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0x458925, 0x458925, 0x458925, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0x458925, 0x458925, 0x458925, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0, 0x656F84, 0x458925, 0x656F84, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0, 0, 0x458925, 0x656F84, 0x656F84, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0, 0, 0x458925, 0x656F84, 0x656F84, 0x656F84, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0, 0, 0x458925, 0x656F84, 0x656F84, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0, 0, 0x458925, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0, 0, 0x458925, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0, 0, 0x458925, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0x458925, 0, 0, 0x458925, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0x458925, 0x458925, 0, 0, 0x458925, 0x458925, 0, 0, 0, 0, 0},
	},
	Align: CenterAlign,
}

// Draws sprites column on screen
func drawSpritesColumn(screen screen.Screen, col int, angle Degree, distanceToWall float64) {
	// We should sort all our sprites by distance so closest are rendered last
	sort.Slice(sprites, func(i, j int) bool {
		return playerDistToSprite(*sprites[i]) > playerDistToSprite(*sprites[j])
	})

	for _, sprite := range sprites {
		drawSpriteColumn(screen, *sprite, col, angle, distanceToWall)
	}
}

// Draws one column for specific sprite on screen
func drawSpriteColumn(screen screen.Screen, sprite Sprite, col int, angle Degree, distanceToWall float64) {
	if playerDistToSprite(sprite) < distanceToWall {
		sees, sP := seesSprite(angle, sprite)
		if !sees {
			return
		}

		spriteHeight := int(float64(distToHeight(playerDistToSprite(sprite), screen.Height())) * sprite.Scale / 2)
		if spriteHeight == 0 {
			spriteHeight = 1 // If scale is set to something really low we could end up with 0 height
		}
		spriteScreenRow := calculateSpriteStart(screen, sprite, spriteHeight)
		spriteCol := int(math.Round(sP * float64(len(sprite.Texture[0])-1)))
		scaledTexture := scaleStringTextureVertically(sprite.Texture, spriteHeight+1)
		for i := 0; i <= spriteHeight; i++ {
			if spriteScreenRow >= 0 {
				if scaledTexture[i][spriteCol] != 0 {
					_ = screen.SetPixel(col, spriteScreenRow, scaledTexture[i][spriteCol])
				}
			}
			spriteScreenRow++
		}
	}
}

// Calculates where on screen a sprite starts rendering (may return negative value, meaning starting point is UP out of screen bounds)
func calculateSpriteStart(screen screen.Screen, sprite Sprite, spriteHeight int) int {
	switch sprite.Align {
	case BottomAlign:
		return screen.Height() / 2
	case CenterAlign:
		return screen.Height()/2 - spriteHeight/2
	case TopAlign:
		return screen.Height()/2 - spriteHeight
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

// Calculates distance to sprite from the given points
func distToSprite(x, y float64, sprite Sprite) float64 {
	return math.Sqrt((sprite.X-x)*(sprite.X-x) + (sprite.Y-y)*(sprite.Y-y))
}

// Calculates if the given point intersects with any sprite
func intersectsWithSprite(x, y float64) (intersects bool) {
	for _, sprite := range sprites {
		if distToSprite(x, y, *sprite) < sprite.Scale/2 {
			return true
		}
	}
	return false
}
