package maths

import "github.com/iCiaran/ray-tracing/maths"

type Solid struct {
	value *maths.Colour
}

func NewSolid(r, g, b float64) *Solid {
	return &Solid{maths.NewVec3(r, g, b)}
}

func (ts *Solid) Value(u, v float64, p *maths.Point3) *maths.Colour {
	return ts.value
}
