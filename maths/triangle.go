package maths

type Triangle struct {
	V      []*Point3
	N      []*Vec3
	Mat    Material
	smooth bool
}

func NewTriangle(va, vb, vc *Point3, na, nb, nc *Vec3, mat Material, smooth bool) *Triangle {
	return &Triangle{[]*Point3{va, vb, vc}, []*Vec3{na, nb, nc}, mat, smooth}
}

func (t *Triangle) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool {
	e1 := Sub(t.V[1], t.V[0])
	e2 := Sub(t.V[2], t.V[0])

	pv := Cross(r.Direction(), e2)
	det := Dot(e1, pv)

	if det < 0.001 {
		return false
	}

	tv := Sub(r.Origin(), t.V[0])
	u := Dot(tv, pv)

	if u < 0.0 || u > det {
		return false
	}

	qv := Cross(tv, e1)

	v := Dot(r.Direction(), qv)
	if v < 0 || (u+v) > det {
		return false
	}

	tr := Dot(e2, qv)

	invDet := 1.0 / det

	tr *= invDet
	u *= invDet
	v *= invDet

	if tr > tMin && tr < tMax {
		rec.T = tr
		rec.P = r.At(rec.T)
		if t.smooth {
			rec.SetFaceNormal(r, t.smoothNormal(u, v))
		} else {
			rec.SetFaceNormal(r, t.flatNormal())
		}
		rec.Mat = t.Mat
		return true
	}

	return false
}

func (t *Triangle) flatNormal() *Vec3 {
	return Add(t.N[0], t.N[1]).Add(t.N[2]).Div(3)
}

func (t *Triangle) smoothNormal(u, v float64) *Vec3 {
	return Mul(t.N[0], 1.0-u-v).Add(Mul(t.N[1], u)).Add(Mul(t.N[2], v))
}
