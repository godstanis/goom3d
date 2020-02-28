package engine

import "math"

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
// For example we have point x:3.2;y:3.9. It's obvious that
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

// Because we use lazy raycasting with steps, we should perform additional rounding to get exact intersection with our grid (i.e. walls)
// todo: reduce complexity and improve readability
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
