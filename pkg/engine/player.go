// Game engine
package engine

// Position of a player relative to a world coordinates
var curX = 0.0
var curY = 0.0

// Current POV of a player. 0 is RIGHT if we look at our map. UP is 270 and so on
var curAngle = Degree{Value: 0.0}

// Current FOV of a camera
var curFov = 90.0

// Maximum raycasting distance
var viewDistance = 20.0

// Turns player around by a given angle (minus is left, plus is right)
func TurnPlayer(dAngle float64) {
	curAngle.Add(dAngle)
}

// Moves player vertically (forward, backward) by a given dist (related to it's angle)
func StrafePlayerV(dDist float64) {
	nextX, nextY := increaseDegreeVector(curX, curY, curAngle, dDist)
	if CollidesWithAnything(nextX, nextY) {
		return
	}
	curX, curY = nextX, nextY
}

// Moves player horizontally (left, right) by a given dist (related to it's angle)
func StrafePlayerH(dDist float64) {
	nextX, nextY := increaseDegreeVector(curX, curY, Degree{curAngle.Plus(90)}, dDist)
	if CollidesWithAnything(nextX, nextY) {
		return
	}
	curX, curY = nextX, nextY
}

// Determines if the given point collides with anny solid object
func CollidesWithAnything(x, y float64) bool {
	intersectsWithMap, _, _ := IntersectsWithMap(x, y)
	if intersectsWithMap || intersectsWithSprite(x, y) {
		return true
	}
	return false
}

// Changes FOV by a given amount
func ShiftFov(fov float64) {
	curFov += fov
}

// Set's player global position relative to world coordinates
func SetPlayerPosition(x, y float64, angle Degree) {
	curX = x
	curY = y
	curAngle = angle
}
