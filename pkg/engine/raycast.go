// Ray casting related stuff
package engine

import (
	"math"
)

// Casts a ray with angle (0-360) and returns (hit status, hit distance (passed distance is default), hit title value (0 default), distance is the max length of a ray
//
// 0 angle is UP, 90 is RIGHT and so on (360 is the same as 0)
// (x0,y0,x1,y1) - vector (x0,y0 is a ray starting point)
//
// tileP is a relation of our wall hit to it's starting plane point (for example if a wall is on X axis and we hit it in the middle, tileP is 0.5)
// 	so 0.0 means we hit the beginning of a wall, 1.0 means we hit the last column of a wall
func rayCast(x0, y0 float64, angle Degree, distance float64) (hit bool, dist float64, tile int, tileP float64) {
	length := 0.0 // Length of hit check
	step := 0.01  // Interval of collision checking
	for length <= distance {
		if hit, tile, tileP := IntersectsWithMap(getVectorEnd(x0, y0, angle, length)); hit {
			return true, length, tile, tileP
		}
		length += step
	}
	return false, distance, 0.0, 0.0
}

// Returns vector's and coordinate (vector is a starting position, angle and distance)
// angle - degrees
func getVectorEnd(x, y float64, angle Degree, length float64) (float64, float64) {
	rads := angle.Get() * (math.Pi / 180)
	return length*math.Cos(rads) + x, length*math.Sin(rads) + y
}

// Returns intersect status and tile value
func IntersectsWithMap(x, y float64) (intersects bool, tile int, tilePoint float64) {
	gridX := int(x)
	gridY := int(y)

	if gridX >= 0 && gridY >= 0 && gridY < len(Map) && gridX < len(Map[0]) && Map[gridY][gridX] >= 1 {
		pointInBox := (x >= float64(gridX)) && (x <= float64(gridX)+1) && (y >= float64(gridY)) && (y <= float64(gridY)+1)
		if pointInBox {
			return true, Map[gridY][gridX], getTilePoint(x, y)
		}
	}

	return false, 0, 0
}

// Determines current tile hit point relative to it's plane (for example 0.3 means we've hit 0/3 point of a wall)
//
// For example we have point X:3.2;Y:3.9. It's obvious that if we stick it to our grid it is 3.2, 4.
func getTilePoint(x, y float64) (tilePoint float64) {
	gridX, gridY := intersectToGrid(x, y)

	pointX := gridX - float64(int(x))
	pointY := gridY - float64(int(y))

	if gridX != x {
		return math.Abs(pointY)
	}
	return math.Abs(pointX)
}

// Because we use lazy raycasting with steps, we should perform additional rounding to get exact intersection with our grid (i.e. walls)
func intersectToGrid(x, y float64) (xi, yi float64) {
	// Calculating distance of our current coordinates to the closest grid coordinates
	distToX := math.Abs(math.Round(x) - x)
	distToY := math.Abs(math.Round(y) - y)

	if distToX < distToY {
		return math.Round(x), y
	}
	return x, math.Round(y)
}
