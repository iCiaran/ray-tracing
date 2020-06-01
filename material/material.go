package maths

import "github.com/iCiaran/ray-tracing/maths"

type Material interface {
	Scatter(rIn *maths.Ray, rec *hittable.HitRecord, attenuation *maths.Colour, scattered *maths.Ray) bool
}
