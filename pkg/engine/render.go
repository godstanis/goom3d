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
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 3, 3, 3, 0, 3, 3, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 2, 2, 2, 0, 2, 2, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 4, 4, 4, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

var TilesSymbol = map[int]string{1: "*", 2: "#", 3: "Z", 4: "+"}

func RenderView(screen *console.Screen) {
	start := time.Now()
	lAngle := curAngle - float64(screen.Width()/2)

	for i := 0; i <= screen.Width(); i++ {
		ok, distance, tile := Trace(curX, curY, lAngle+float64(i), viewDistance)

		if ok {
			DrawWall(screen, tile, i, DistToHeight(distance, screen.Height()))
		}
	}

	console.Render(screen, fmt.Sprintf("FPS: %f; view_angle:%f; player_pos:(x:%f,y:%f)", 1/time.Since(start).Seconds(), curAngle, curX, curY))
}

func DistToHeight(dist float64, screenHeight int) int {
	height := int(float64(screenHeight) / dist)
	if height <= 0 {
		return 1
	}
	return height
}

// Draw a wall on i vertical row of screen
func DrawWall(screen *console.Screen, tile int, i int, height int) {
	var offset int
	if height > screen.Height() {
		offset = 0
	} else {
		offset = (screen.Height() - height) / 2
	}
	for j := offset; j < screen.Height()-offset-1; j++ {
		symbol := TilesSymbol[tile]
		screen.SetPixel(i, j, symbol)
	}
}

// Traces a ray with angle (0-360) and returns (hit status, hit distance (0 default), hit title value (0 default), distance is the max length of a ray
// 0 angle is UP, 90 is RIGHT and so on (360 is the same as 0)
// (x0,y0,x1,y1) - vector (x0,y0 is a camera position)
func Trace(x0, y0, angle float64, distance float64) (bool, float64, int) {
	length := 0.0
	step := 0.01 // Interval of collision checking
	angle += 270 // So that our 0 angle is UP on our Map
	for length < distance {
		if ok, tile := IntersectsWithMap(GetVectorEnd(x0, y0, angle, length)); ok {
			return true, length, tile
		}
		length += step
	}
	return false, 0.0, 0.0
}

// angle - degrees
func GetVectorEnd(x, y float64, angle float64, length float64) (float64, float64) {
	rads := angle * (math.Pi / 180)
	return length*math.Cos(rads) + x, length*math.Sin(rads) + y
}

// Returns intersect status and tile value
func IntersectsWithMap(x, y float64) (bool, int) {
	xStart := int(x)
	yStart := int(y)

	if yStart < len(Map) && xStart < len(Map[0]) {
		pointInBox := (x >= float64(xStart)) && (x <= float64(xStart)+1) && (y >= float64(yStart)) && (y <= float64(yStart)+1)
		if pointInBox && Map[yStart][xStart] >= 1 {
			return true, Map[yStart][xStart]
		}
	}

	return false, 0
}
