package engine

import (
	"glfun/pkg/screen"
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
	Texture [][]string
}

// sprites information for map
var sprites = []*Sprite{
	&lampSprite,
	&boxSprite,
	&box2Sprite,
	&paintingSprite,
	&tableSprite,
	&chairSprite,
}

var lampSprite = Sprite{
	X: 1.4, Y: 2.5, Texture: [][]string{
		{"X", "X", "X", "-", "-", "X", "X", "X"},
		{"X", "X", "X", "|", "|", "X", "X", "X"},
		{"O", "X", "X", "%", "%", "X", "X", "O"},
		{"%", "%", "X", "%", "%", "X", "%", "%"},
		{"X", "X", "%", "%", "%", "%", "X", "X"},
		{"X", "X", "X", "%", "%", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
	},
	Align: TopAlign,
}

var boxSprite = Sprite{
	X: 1.2, Y: 1.2, Texture: [][]string{
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "=", "=", "=", "=", "X", "X"},
		{"X", "|", "-", "-", "-", "-", "|", "X"},
		{"|", "-", "-", "-", "-", "-", "-", "|"},
		{"|", "-", "-", "-", "-", "-", "-", "|"},
		{"X", "|", "-", "-", "-", "-", "|", "X"},
		{"X", "X", "=", "=", "=", "=", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
	},
	Align: BottomAlign,
}

var box2Sprite = Sprite{
	X: 1.5, Y: 1.7, Texture: [][]string{
		{"X", "X", "X", "X", "X", "X", "X", "X"},
		{"X", "X", "=", "=", "=", "=", "X", "X"},
		{"X", "|", "-", "-", "-", "-", "|", "X"},
		{"|", "-", "-", "-", "-", "-", "-", "|"},
		{"|", "-", "-", "-", "-", "-", "-", "|"},
		{"X", "|", "-", "-", "-", "-", "|", "X"},
		{"X", "X", "=", "=", "=", "=", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X", "X"},
	},
	Align: BottomAlign,
}

var tableSprite = Sprite{
	X: 1.4, Y: 2.5, Texture: [][]string{
		{"X", "X", "X", "X", "X", "X", "X"},
		{"X", "=", "=", "=", "=", "=", "="},
		{"=", "=", "=", "=", "=", "=", "|"},
		{"|", "X", "|", "X", "|", "X", "|"},
		{"|", "X", "|", "X", "|", "X", "|"},
		{"|", "X", "|", "X", "|", "X", "|"},
		{"|", "X", "X", "X", "|", "X", "X"},
		{"X", "X", "X", "X", "X", "X", "X"},
	},
	Align: BottomAlign,
}

var chairSprite = Sprite{
	X: 1.4, Y: 2.1, Texture: [][]string{
		{"X", "X", "X", "-", "X"},
		{"X", "X", "X", "|", "X"},
		{"X", "X", "X", "|", "X"},
		{"=", "=", "=", "=", "X"},
		{"|", "X", "X", "|", "X"},
		{"|", "X", "X", "|", "X"},
		{"X", "X", "X", "X", "X"},
		{"X", "X", "X", "X", "X"},
	},
	Align: BottomAlign,
}

var paintingSprite = Sprite{
	X: 2.55, Y: 1.41, Texture: [][]string{
		{"#", "-", "-", "-", "-", "-", "-", "-", "-", "-", "#"},
		{"|", "▒", "▒", "░", "░", "░", "▓", "▓", "▓", "░", "|"},
		{"|", "▒", "▒", "░", "░", "░", "▓", "▓", "░", "░", "|"},
		{"|", "▒", "░", "░", "░", "░", "░", "▓", "░", "▓", "|"},
		{"|", "▒", "░", "░", "░", "░", "▓", "▓", "▓", "░", "|"},
		{"|", "░", "░", "░", "░", "▓", "░", "▓", "░", "░", "|"},
		{"|", "░", "░", "░", "░", "░", "░", "▓", "░", "░", "|"},
		{"|", "░", "░", "░", "░", "░", "▓", "░", "▓", "░", "|"},
		{"|", "░", "░", "░", "░", "░", "▓", "░", "▓", "░", "|"},
		{"|", "░", "░", "░", "░", "░", "▓", "░", "▓", "░", "|"},
		{"#", "-", "-", "-", "-", "-", "-", "-", "-", "-", "#"},
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
		sees, sP := seesSprite(screen, angle, sprite)
		if !sees {
			return
		}

		spriteHeight := distToHeight(playerDistToSprite(sprite), screen.Height()) / 2
		spriteScreenRow := calculateSpriteStart(screen, sprite, spriteHeight)
		spriteCol := int(math.Round(sP * float64(len(sprite.Texture[0])-1)))

		for i := 0; i <= spriteHeight; i++ {
			if spriteScreenRow >= 0 {
				spriteRow := int(math.Round((float64(i) / float64(spriteHeight)) * float64(len(sprite.Texture)-1)))
				if sprite.Texture[spriteRow][spriteCol] != "X" {
					_ = screen.SetPixel(col, spriteScreenRow, sprite.Texture[spriteRow][spriteCol])
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
func seesSprite(screen screen.Screen, angle Degree, sprite Sprite) (hit bool, spritePoint float64) {
	angleDiff := float64(int(angle.Get()-playerAngleToSprite(sprite).Get()+180+360)%360 - 180)
	angleOffset := (float64(screen.Width()) / 10) / playerDistToSprite(sprite) / 2 // todo: investigate skewing artifacts on resolution changes

	if !(angleDiff >= -angleOffset && angleDiff <= angleOffset) {
		return false, 0
	}
	// Calculate point relation to hit
	return true, (angleDiff + angleOffset) / (angleOffset * 2)
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
	return math.Sqrt((sprite.X-curX)*(sprite.X-curX) + (sprite.Y-curY)*(sprite.Y-curY))
}
