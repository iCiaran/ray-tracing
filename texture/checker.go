package maths

import (
	"math"

	"github.com/iCiaran/ray-tracing/maths"
)

type Checker struct {
	odd   Texture
	even  Texture
	width float64
}

func NewChecker(t0, t1 Texture, w float64) *Checker {
	return &Checker{t0, t1, 1.0 / w}
}

func (tc *Checker) Value(u, v float64, p *maths.Vec3) *maths.Colour {
	sines := math.Sin(tc.width*p.X()) * math.Sin(tc.width*p.Y()) * math.Sin(tc.width*p.Z())
	if sines < 0 {
		return tc.odd.Value(u, v, p)
	}
	return tc.even.Value(u, v, p)
}
