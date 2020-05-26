package maths

type AABB struct {
	min *Point3
	max *Point3
}

func NewAABB(a, b *Point3) *AABB {
	return &AABB{a, b}
}

func (a *AABB) Hit(r *Ray, tMin, tMax float64) bool {
	for i := 0; i < 3; i++ {
		invD := 1.0 / r.Direction()[i]
		t0 := invD * (a.min[i] - r.Origin()[i])
		t1 := invD * (a.max[i] - r.Origin()[i])

		if invD < 0.0 {
			t0, t1 = t1, t0
		}

		if t0 > tMin {
			tMin = t0
		}

		if t1 < tMax {
			tMax = t1
		}

		if tMax <= tMin {
			return false
		}
	}
	return true
}
