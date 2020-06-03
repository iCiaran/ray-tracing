package hittable

import (
	"math"

	"github.com/iCiaran/ray-tracing/hittable/hitmatrecord"
	"github.com/iCiaran/ray-tracing/material"
	"github.com/iCiaran/ray-tracing/maths"
)

type Triangle struct {
	V      []*maths.Point3
	N      []*maths.Vec3
	T      []*maths.Point3
	Mat    material.Material
	smooth bool
}

func NewTriangle(va, vb, vc, na, nb, nc, ta, tb, tc *maths.Vec3, mat material.Material, smooth bool) *Triangle {
	return &Triangle{[]*maths.Point3{va, vb, vc}, []*maths.Vec3{na, nb, nc}, []*maths.Point3{ta, tb, tc}, mat, smooth}
}

func (t *Triangle) Hit(r *maths.Ray, tMin, tMax float64, rec *hitmatrecord.HitMatRecord) bool {
	e1 := maths.Sub(t.V[1], t.V[0])
	e2 := maths.Sub(t.V[2], t.V[0])

	pv := maths.Cross(r.Direction(), e2)
	det := maths.Dot(e1, pv)

	if det < 0.001 {
		return false
	}

	tv := maths.Sub(r.Origin(), t.V[0])
	u := maths.Dot(tv, pv)

	if u < 0.0 || u > det {
		return false
	}

	qv := maths.Cross(tv, e1)

	v := maths.Dot(r.Direction(), qv)
	if v < 0 || (u+v) > det {
		return false
	}

	tr := maths.Dot(e2, qv)

	invDet := 1.0 / det

	tr *= invDet
	u *= invDet
	v *= invDet

	if tr > tMin && tr < tMax {
		rec.Rec.T = tr
		rec.Rec.P = r.At(rec.Rec.T)
		if t.smooth {
			rec.Rec.SetFaceNormal(r, interpolate(u, v, t.N[0], t.N[1], t.N[2]).Normalise())
		} else {
			rec.Rec.SetFaceNormal(r, maths.Cross(e1, e2).Normalise())
		}
		texUV := interpolate(u, v, t.T[0], t.T[1], t.T[2])
		rec.Rec.U = texUV.X()
		rec.Rec.V = texUV.Y()
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

	*outputBox = *NewAABB(maths.NewVec3(smallX, smallY, smallZ), maths.NewVec3(bigX, bigY, bigZ))
	return true
}

func interpolate(u, v float64, a, b, c *maths.Vec3) *maths.Vec3 {
	return maths.Mul(a, 1.0-u-v).Add(maths.Mul(b, u)).Add(maths.Mul(c, v))
}
