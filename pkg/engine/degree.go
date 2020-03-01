package engine

// Degree represents consistent 0-360 degree service object
type Degree struct {
	Value float64
}

// Creates new angle object
func (dgr Degree) NewDegree(angle float64) Degree {
	ndgr := Degree{}
	ndgr.Set(angle)
	return ndgr
}

// Sets and normalizes angle
func (dgr *Degree) Set(angle float64) {
	if angle >=360 {
		angle = angle - 360
	}
	if angle < 0 {
		angle = 360 + angle
	}
	dgr.Value = angle
}

// Adds angle
func (dgr *Degree) Add(angle float64) {
	dgr.Set(dgr.Get()+angle)
}

// Returns adjusted angle value, no sets performed
func (dgr Degree) Plus(angle float64) float64 {
	dgr.Set(dgr.Get()+angle)
	return dgr.Get()
}

// Gets stored angle value
func (dgr Degree) Get() float64 {
	return dgr.Value
}