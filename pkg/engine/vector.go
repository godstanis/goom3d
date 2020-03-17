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
func (v Vector) Angle() float64 {
	angle := math.Atan2(v.Y, v.X) * (180 / math.Pi)
	if angle < 0 {
		return 360 + angle
	}
	return angle
}

// Magnitude calculates vector scalar length
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
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
