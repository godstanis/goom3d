package engine

// Degree represents consistent 0-360 degree service object
type Degree struct {
	Value float64
}

// NewDegree creates new angle object
func (dgr Degree) NewDegree(angle float64) Degree {
	ndgr := Degree{}
	ndgr.Set(angle)
	return ndgr
}

// Set sets and normalizes angle
func (dgr *Degree) Set(angle float64) {
	if angle >= 360 {
		angle = angle - 360
	}
	if angle < 0 {
		angle = 360 + angle
	}
	dgr.Value = angle
}

// Add adds angle
func (dgr *Degree) Add(angle float64) {
	dgr.Set(dgr.Get() + angle)
}

// Plus returns adjusted angle value, no sets performed
func (dgr Degree) Plus(angle float64) float64 {
	dgr.Set(dgr.Get() + angle)
	return dgr.Get()
}

// Get gets stored angle value
func (dgr Degree) Get() float64 {
	return dgr.Value
}
