package maths

import "math"

type AABB struct {
	min *Point3
	max *Point3
}

func NewAABB(a, b *Point3) *AABB {
	return &AABB{a, b}
}

func (a *AABB) Hit(r *Ray, tMin, tMax float64) bool {
	for i := 0; i < 3; i++ {
		t0 := math.Min(a.min[i]-r.Origin()[i]/r.Direction()[i],
			a.max[i]-r.Origin()[i]/r.Direction()[i])
		t1 := math.Max(a.min[i]-r.Origin()[i]/r.Direction()[i],
			a.max[i]-r.Origin()[i]/r.Direction()[i])
		tMin = math.Max(t0, tMin)
		tMax = math.Min(t1, tMax)
		if tMax <= tMin {
			return false
		}
	}
	return true
}
