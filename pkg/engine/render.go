package engine

import (
	"fmt"
	"glfun/pkg/console"
	"math"
	"time"
)

// Map structure
//
// 0 - empty space
// 1 - solid wall
var Map = make([][]int, 1)

var WorldMap = [][]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 3, 4, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 0, 1},
	{1, 0, 2, 0, 2, 0, 0, 0, 0, 2, 2, 0, 0, 1},
	{1, 2, 2, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

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
		{"▒", "░", "□", "□", "□", "░", "▓"},
		{"▒", "░", "□", "□", "□", "░", "▒"},
		{"▒", "░", "░", "░", "░", "░", "▒"},
		{"▒", "░", "⨀", "░", "░", "░", "▒"},
		{"▒", "░", "░", "░", "░", "░", "▓"},
		{"▒", "░", "░", "░", "░", "░", "▒"},
		{"▒", "▒", "▒", "▒", "▒", "▒", "▒"},
	},
}

// Renders a frame
func RenderView(screen *console.Screen) {
	start := time.Now()
	lAngle := curAngle - curFov/2

	for i := 0; i <= screen.Width(); i++ {
		traversed := float64(i) / float64(screen.Width()) // How much of a screen space has been traversed (0.0 to 1.0, i.e. 0.3 is 30%)
		hit, distance, tile, tileP := RayCast(curX, curY, lAngle+(curFov*traversed), viewDistance)

		if hit {
			DrawTexturedWall(screen, tile, i, DistToHeight(distance, screen.Height()), tileP)
		}
	}

	console.Render(screen, fmt.Sprintf("FPS: %6.4f; POV:%4.2f; FOV:%4.2f, player_pos:(x:%9.4f,y:%9.4f)", 1/time.Since(start).Seconds(), curAngle, curFov, curX, curY))
}

// Draws textured wall column on screen
func DrawTexturedWall(screen *console.Screen, tile int, i int, height int, tileP float64) {
	var offset int
	if height > screen.Height() {
		offset = 0
	} else {
		offset = (screen.Height() - height) / 2
	}
	scaledTextureRow := ScaleTextureV(tile, height, tileP)
	for j := offset; j < screen.Height()-offset-1; j++ {
		screen.SetPixel(i, j, scaledTextureRow[j-offset])
	}
}

// Determines final height of a wall column on screen
func DistToHeight(dist float64, screenHeight int) int {
	height := int(float64(screenHeight) / dist)
	if height < 0 {
		return 0
	}
	return height
}

// Scales texture column to h proportionally
func ScaleTextureV(tile, h int, tileP float64) []string {
	c := TranslateToTextureCol(tile, tileP)
	var textureCol, scaled []string
	for _, val := range TileTextures[tile] {
		textureCol = append(textureCol, val[c])
	}
	if len(textureCol)-1 == h-1 {
		return textureCol // No need for scaling if required height is equal to texture's original height
	}
	perTile := float64(len(textureCol)-1) / float64(h-1)
	curTileS := 0.0
	for i := 0; i < h; i++ {
		scaled = append(scaled, TileTextures[tile][int(math.Round(curTileS))][c])
		curTileS += perTile
	}

	if len(scaled) > 0 && len(textureCol) > 0 {
		scaled[0] = textureCol[0]
		scaled[len(scaled)-1] = textureCol[len(textureCol)-1]
	}

	return scaled
}

// Translates percentages to concrete texture column
func TranslateToTextureCol(tile int, tileP float64) int {
	textureWidth := len(TileTextures[tile][0])
	return int(math.Round(float64(textureWidth-1) * tileP))
}

// Casts a ray with angle (0-360) and returns (hit status, hit distance (0 default), hit title value (0 default), distance is the max length of a ray
// 0 angle is UP, 90 is RIGHT and so on (360 is the same as 0)
// (x0,y0,x1,y1) - vector (x0,y0 is a camera position)
func RayCast(x0, y0, angle float64, distance float64) (hit bool, tile float64, dist int, tileP float64) {
	length := 0.0
	step := 0.01 // Interval of collision checking
	angle += 270 // So that our 0 angle is UP on our Map
	for length < distance {
		if ok, tile, tileP := IntersectsWithMap(GetVectorEnd(x0, y0, angle, length)); ok {
			return true, length, tile, tileP
		}
		length += step
	}
	return false, 0.0, 0.0, 0.0
}

// angle - degrees
func GetVectorEnd(x, y float64, angle float64, length float64) (float64, float64) {
	rads := angle * (math.Pi / 180)
	return length*math.Cos(rads) + x, length*math.Sin(rads) + y
}

// Returns intersect status and tile value
func IntersectsWithMap(x, y float64) (intersects bool, tile int, tilePoint float64) {
	xStart := int(x)
	yStart := int(y)

	if yStart < len(Map) && xStart < len(Map[0]) {
		pointInBox := (x >= float64(xStart)) && (x <= float64(xStart)+1) && (y >= float64(yStart)) && (y <= float64(yStart)+1)
		if pointInBox && Map[yStart][xStart] >= 1 {
			return true, Map[yStart][xStart], TilePoint(x, y)
		}
	}

	return false, 0, 0
}

// Determines current tile hit point relative to it's plane (for example 0.3 means we've hit 0/3 point of a wall)
func TilePoint(x, y float64) (tilePoint float64) {
	flatX, flatY := FlattenIntersect(x, y)
	pointX := flatX - float64(int(x))
	pointY := flatY - float64(int(y))

	if pointX == 0.0 {
		return pointY
	} else {
		return pointX
	}
}

//todo: reduce complexity and improve readability
func FlattenIntersect(x, y float64) (xi, yi float64) {
	tileXf, tileYf := float64(int(x)), float64(int(y))

	distToStartX, distToEndX := x-tileXf, tileXf+1-x
	distToStartY, distToEndY := y-tileYf, tileYf+1-y

	closeToStartX := distToStartX < distToEndX
	closeToStartY := distToStartY < distToEndY

	var distToX, distToY float64

	if closeToStartX {
		distToX = distToStartX
	} else {
		distToX = distToEndX
	}
	if closeToStartY {
		distToY = distToStartY
	} else {
		distToY = distToEndY
	}

	baseIsX := distToX < distToY

	var resX, resY float64

	if baseIsX {
		resX = math.Round(x)
		resY = y
	} else {
		resX = x
		resY = math.Round(y)
	}

	return resX, resY
}
