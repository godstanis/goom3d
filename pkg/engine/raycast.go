package engine

import (
	"math"
)

// Casts a ray using vector direction and returns (hit status, hit distance (passed distance is default), hit title value (0 default), distance is the max length of a ray
//
// tileP is a relation of our wall hit to it's starting plane point (for example if a wall is on X axis and we hit it in the middle, tileP is 0.5)
// 	so 0.0 means we hit the beginning of a wall, 1.0 means we hit the last column of a wall
func rayCast(x0, y0 float64, dir Vector, distance float64) (hit bool, dist float64, tile int, tileP float64) {
	length := 0.0 // Length of hit check
	step := 0.01  // Interval of collision checking
	for length <= distance {
		if hit, tile, tileP := IntersectsWithMap(x0, y0); hit {
			return true, length, tile, tileP
		}
		x0, y0 = x0+dir.X*step, y0+dir.Y*step
		length += step
	}
	return false, distance, 0.0, 0.0
}

// IntersectsWithMap returns intersect status and tile value
func IntersectsWithMap(x, y float64) (intersects bool, tile int, tilePoint float64) {
	gridX, gridY := math.Floor(x), math.Floor(y)

	if gridX >= 0 && gridY >= 0 && gridY < float64(len(Map)) && gridX < float64(len(Map[0])) {
		if (x >= gridX) && (x <= gridX+1) && (y >= gridY) && (y <= gridY+1) { // if point is in square boundaries
			if tile := Map[int(gridY)][int(gridX)]; tile > 0 {
				gX, gY := intersectToGrid(x, y)   // Round x or y to closest axis
				if gX == gridX+1 || gY == gridY { // We should inverse texture for this case
					return true, tile, 1 - getTilePoint(gX, gY)
				}
				return true, tile, getTilePoint(gX, gY)
			}
		}
	}

	return false, 0, 0
}

// Determines current tile hit point relative to it's plane (for example 0.3 means we've hit 0/3 point of a wall)
//
// For example we have point X:3.2;Y:3.9. It's obvious that if we stick it to our grid it is 3.2, 4.
func getTilePoint(gridX, gridY float64) (tilePoint float64) {
	baseX := float64(int(gridX))
	if gridX == baseX {
		return math.Abs(gridY - float64(int(gridY)))
	}
	return math.Abs(gridX - baseX)
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
