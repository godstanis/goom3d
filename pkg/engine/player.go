// Game engine
package engine

import "math"

// Position of a player relative to a world coordinates
var curX = 0.0
var curY = 0.0

// Current POV of a player
var curAngle = 0.0

// Current FOV of a camera
var curFov = 90.0

// Maximum raycasting distance
var viewDistance = 20.0

// Turns player around by a given angle (minus is left, plus is right)
func TurnPlayer(dAngle float64) {
	curAngle += dAngle
	
	if curAngle > 360 {
		curAngle = 0 + math.Abs(dAngle)
	}
	if curAngle < 0 {
		curAngle = 360 - math.Abs(dAngle)
	}
}

// Moves player vertically (forward, backward) by a given dist (related to it's angle)
func StrafePlayerV(dDist float64) {
	curX, curY = GetVectorEnd(curX, curY, curAngle-90, dDist)
}

// Moves player horizontally (left, right) by a given dist (related to it's angle)
func StrafePlayerH(dDist float64) {
	curX, curY = GetVectorEnd(curX, curY, curAngle, dDist)
}

// Changes FOV by a given amount
func ShiftFov(fov float64) {
	curFov += fov
}

// Set's player global position relative to world coordinates
func SetPlayerPosition(x, y, angle float64) {
	curX = x
	curY = y
	curAngle = angle
}
