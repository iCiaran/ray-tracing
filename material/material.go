package material

import (
	"github.com/iCiaran/ray-tracing/hittable/hitrecord"
	"github.com/iCiaran/ray-tracing/maths"
)

type Material interface {
	Scatter(rIn *maths.Ray, rec *hitrecord.HitRecord, attenuation *maths.Colour, scattered *maths.Ray) bool
}
