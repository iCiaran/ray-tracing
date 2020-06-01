package maths

import (
	"math"
)

type Triangle struct {
	V      []*Point3
	N      []*Vec3
	T      []*Point3
	Mat    Material
	smooth bool
}

func NewTriangle(va, vb, vc *Point3, na, nb, nc *Vec3, ta, tb, tc *Vec3, mat Material, smooth bool) *Triangle {
	return &Triangle{[]*Point3{va, vb, vc}, []*Vec3{na, nb, nc}, []*Point3{ta, tb, tc}, mat, smooth}
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
			rec.SetFaceNormal(r, interpolate(u, v, t.N[0], t.N[1], t.N[2]).Normalise())
		} else {
			rec.SetFaceNormal(r, Cross(e1, e2).Normalise())
		}
		texUV := interpolate(u, v, t.T[0], t.T[1], t.T[2])
		rec.U = texUV.X()
		rec.V = texUV.Y()
		rec.Mat = t.Mat
		return true
	}

	return false
}

func (t *Triangle) BoundingBox(t0, t1 float64, outputBox *AABB) bool {
	smallX := math.Min(math.Min(t.V[0].X(), t.V[1].X()), t.V[2].X())
	smallY := math.Min(math.Min(t.V[0].Y(), t.V[1].Y()), t.V[2].Y())
	smallZ := math.Min(math.Min(t.V[0].Z(), t.V[1].Z()), t.V[2].Z())
	bigX := math.Max(math.Max(t.V[0].X(), t.V[1].X()), t.V[2].X())
	bigY := math.Max(math.Max(t.V[0].Y(), t.V[1].Y()), t.V[2].Y())
	bigZ := math.Max(math.Max(t.V[0].Z(), t.V[1].Z()), t.V[2].Z())

	*outputBox = *NewAABB(NewVec3(smallX, smallY, smallZ), NewVec3(bigX, bigY, bigZ))
	return true
}

func interpolate(u, v float64, a, b, c *Vec3) *Vec3 {
	return Mul(a, 1.0-u-v).Add(Mul(b, u)).Add(Mul(c, v))
}
