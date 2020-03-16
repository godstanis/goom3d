package engine

import "math"

// Vector represents 2d vector
type Vector struct {
	X, Y float64
}

// NewFromAngle returns a unit vector for a given angle (relative to positive X-axis)
func (v Vector) NewFromAngle(angle float64) Vector {
	rads := angle * (math.Pi / 180)
	dx, dy := math.Cos(rads), math.Sin(rads)
	return Vector{X: dx, Y: dy}
}

// Angle transforms vector to angle (relative to positive X-axis)
func (v Vector) Angle() Degree {
	angle := math.Atan2(v.Y, v.X) * (180 / math.Pi)
	return Degree{}.NewDegree(angle)
}

// Rotate rotates vector
func (v *Vector) Rotate(angle float64) {
	rads := angle * (math.Pi / 180)
	nX, nY := v.X*math.Cos(rads)-v.Y*math.Sin(rads), v.X*math.Sin(rads)+v.Y*math.Cos(rads)
	v.X, v.Y = nX, nY
}

// NewRotated creates a new rotated vector
func (v Vector) NewRotated(angle float64) Vector {
	v.Rotate(angle)
	return Vector{X: v.X, Y: v.Y}
}

// Position of a player relative to a world coordinates
var curX = 0.0
var curY = 0.0

// Current POV of a player direction vector
var curVector = Vector{}

// Current FOV of a camera
var curFov = 90.0

// Maximum raycasting distance
var viewDistance = 20.0

// TurnPlayer turns player around by a given angle (minus is left, plus is right)
func TurnPlayer(dAngle float64) {
	curVector.Rotate(dAngle)
}

// StrafePlayerV moves player vertically (forward, backward) by a given dist (related to it's angle)
func StrafePlayerV(dDist float64) {
	nextX, nextY := curX+curVector.X*dDist, curY+curVector.Y*dDist
	if CollidesWithAnything(nextX, nextY) {
		return
	}
	curX, curY = nextX, nextY
}

// StrafePlayerH moves player horizontally (left, right) by a given dist (related to it's angle)
func StrafePlayerH(dDist float64) {
	nextX, nextY := curX-curVector.Y*dDist, curY+curVector.X*dDist
	if CollidesWithAnything(nextX, nextY) {
		return
	}
	curX, curY = nextX, nextY
}

// CollidesWithAnything determines if the given point collides with anny solid object
func CollidesWithAnything(x, y float64) bool {
	intersectsWithMap, _, _ := IntersectsWithMap(x, y)
	if intersectsWithMap || intersectsWithSprite(x, y) {
		return true
	}
	return false
}

// ShiftFov changes FOV by a given amount
func ShiftFov(fov float64) {
	curFov += fov
}

// SetPlayerPosition sets player global position relative to world coordinates
func SetPlayerPosition(x, y, angle float64) {
	curX = x
	curY = y
	curVector = Vector{}.NewFromAngle(angle)
}
