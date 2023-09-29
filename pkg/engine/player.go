package engine

type Camera struct {
	X, Y     float64
	POV      *Vector
	FOV      float64
	ViewDist float64
}

var player = *NewPlayer()

func NewPlayer() *Camera {
	return &Camera{
		X:        0,
		Y:        0,
		POV:      &Vector{},
		FOV:      90,
		ViewDist: 20,
	}
}

// TurnPlayer turns player around by a given angle (minus is left, plus is right)
func (p *Camera) TurnPlayer(dAngle float64) {
	p.POV.Rotate(dAngle)
}

// StrafePlayerV moves player vertically (forward, backward) by a given dist (related to it's angle)
func (p *Camera) StrafePlayerV(dDist float64) {
	nextX, nextY := p.X+p.POV.X*dDist, p.Y+p.POV.Y*dDist
	if p.Collision(nextX, nextY) {
		return
	}
	p.X, p.Y = nextX, nextY
}

// StrafePlayerH moves player horizontally (left, right) by a given dist (related to it's angle)
func (p *Camera) StrafePlayerH(dDist float64) {
	nextX, nextY := p.X-p.POV.Y*dDist, p.Y+p.POV.X*dDist
	if p.Collision(nextX, nextY) {
		return
	}
	p.X, p.Y = nextX, nextY
}

// Collision determines if the given point collides with anny solid object
func (p *Camera) Collision(x, y float64) bool {
	intersectsWithMap, _, _ := IntersectsWithMap(x, y)
	if intersectsWithMap || intersectsWithSprite(x, y) {
		return true
	}
	return false
}

// ShiftFov changes FOV by a given amount
func (p *Camera) ShiftFov(fov float64) {
	p.FOV += fov
}

// Position sets player global position relative to world coordinates
func (p *Camera) Position(x, y, angle float64) {
	p.X, p.Y = x, y
	vec := (&Vector{}).NewFromAngle(angle)
	p.POV = &vec
}

// TurnPlayer turns player around by a given angle (minus is left, plus is right)
func TurnPlayer(dAngle float64) {
	player.TurnPlayer(dAngle)
}

// StrafePlayerV moves player vertically (forward, backward) by a given dist (related to it's angle)
func StrafePlayerV(dDist float64) {
	player.StrafePlayerV(dDist)
}

// StrafePlayerH moves player horizontally (left, right) by a given dist (related to it's angle)
func StrafePlayerH(dDist float64) {
	player.StrafePlayerH(dDist)
}
