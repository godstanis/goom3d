// Game engine
package engine

var curX = 4.6
var curY = 7.0
var curAngle = 0.0
var curFov = 90.0
var viewDistance = 15.0

func TurnPlayer(dAngle float64) {
	curAngle += dAngle
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
