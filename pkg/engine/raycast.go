package engine

import (
	"math"
)

// Casts a ray with angle (0-360) and returns (hit status, hit distance (0 default), hit title value (0 default), distance is the max length of a ray
// 0 angle is UP, 90 is RIGHT and so on (360 is the same as 0)
// (x0,y0,x1,y1) - vector (x0,y0 is a camera position)
func RayCast(x0, y0, angle float64, distance float64) (hit bool, dist float64, tile int, tileP float64) {
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

// Returns vector's and coordinate (vector is a starting position, angle and distance)
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
//
// For example we have point x:3.2;y:3.9. It's obvious that if we stick it to our grid it is 3.2, 4.
func TilePoint(x, y float64) (tilePoint float64) {
	gridX, gridY := IntersectToGrid(x, y)

	pointX := gridX - math.Round(x)
	pointY := gridY - math.Round(y)

	if pointX == 0.0 {
		return math.Abs(pointY)
	}
	return math.Abs(pointX)
}

// Because we use lazy raycasting with steps, we should perform additional rounding to get exact intersection with our grid (i.e. walls)
func IntersectToGrid(x, y float64) (xi, yi float64) {
	// Calculating distance of our current coordinates to the closest grid coordinates
	distToX := math.Abs(math.Round(x) - x)
	distToY := math.Abs(math.Round(y) - y)

	if distToX < distToY {
		return math.Round(x), y
	}
	return x, math.Round(y)
}
