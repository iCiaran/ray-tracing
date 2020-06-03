package texture

import "github.com/iCiaran/ray-tracing/maths"

type Texture interface {
	Value(u, v float64, p *maths.Point3) *maths.Colour
}
