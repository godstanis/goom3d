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
	// We dont want to calculate sin/cos for each hit point so we only prepare a simplified vector for increaseVector func
	sX, sY := increaseDegreeVector(x0, y0, angle, -1)
	for length <= distance {
		if hit, tile, tileP := IntersectsWithMap(increaseVector(sX, sY, x0, y0, length)); hit {
			return true, length, tile, tileP
		}
		length += step
	}
	return false, distance, 0.0, 0.0
}

// Gets next point for vector by length
func increaseVector(x0, y0, x1, y1, length float64) (x, y float64) {
	dX, dY := x1-x0, y1-y0
	mg := math.Sqrt(dX*dX + dY*dY)
	uX, uY := dX/mg, dY/mg

	return x1 + uX*length, y1 + uY*length
}

// Returns vector's and coordinate (vector is a starting position, angle and distance)
// angle - degrees
func increaseDegreeVector(x, y float64, angle Degree, length float64) (float64, float64) {
	rads := angle.Get() * (math.Pi / 180)
	return length*math.Cos(rads) + x, length*math.Sin(rads) + y
}

// Returns intersect status and tile value
func IntersectsWithMap(x, y float64) (intersects bool, tile int, tilePoint float64) {
	gridX, gridY := math.Floor(x), math.Floor(y)

	if gridX >= 0 && gridY >= 0 && gridY < float64(len(Map)) && gridX < float64(len(Map[0])) {
		if (x >= gridX) && (x <= gridX+1) && (y >= gridY) && (y <= gridY+1) { // if point is in square boundaries
			if tile := Map[int(gridY)][int(gridX)]; tile > 0 {
				gX, gY := intersectToGrid(x, y) // Round x or y to closest axis
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
