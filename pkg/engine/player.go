// Game engine
package engine

import "math"

var curX = 0.0
var curY = 0.0
var curAngle = 0.0
var curFov = 90.0
var viewDistance = 20.0

func TurnPlayer(dAngle float64) {
	curAngle += dAngle
	
	if curAngle > 360 {
		curAngle = 0 + math.Abs(dAngle)
	}
	if curAngle < 0 {
		curAngle = 360 - math.Abs(dAngle)
	}
}

func StrafePlayerV(dDist float64) {
	curX, curY = GetVectorEnd(curX, curY, curAngle-90, dDist)
}

func StrafePlayerH(dDist float64) {
	curX, curY = GetVectorEnd(curX, curY, curAngle, dDist)
}

func ShiftFov(fov float64) {
	curFov += fov
}

func SetPlayerPosition(x, y, angle float64) {
	curX = x
	curY = y
	curAngle = angle
}
