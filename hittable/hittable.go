package hittable

import (
	"github.com/iCiaran/ray-tracing/hittable/hitmatrecord"
	"github.com/iCiaran/ray-tracing/maths"
)

type Hittable interface {
	Hit(r *maths.Ray, tMin, tMax float64, rec *hitmatrecord.HitMatRecord) bool
	BoundingBox(t0, t1 float64, outputBox *AABB) bool
}
