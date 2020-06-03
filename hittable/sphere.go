package hittable

import (
	"math"

	"github.com/iCiaran/ray-tracing/hittable/hitmatrecord"
	"github.com/iCiaran/ray-tracing/material"
	"github.com/iCiaran/ray-tracing/maths"
)

type Sphere struct {
	Center *maths.Point3
	Radius float64
	Mat    material.Material
}

func NewSphere(center *maths.Point3, radius float64, mat material.Material) *Sphere {
	return &Sphere{center, radius, mat}
}

func (s *Sphere) Hit(r *maths.Ray, tMin, tMax float64, rec *hitmatrecord.HitMatRecord) bool {
	oc := maths.Sub(r.Origin(), s.Center)
	a := r.Direction().LenSquared()
	halfB := maths.Dot(oc, r.Direction())
	c := oc.LenSquared() - s.Radius*s.Radius
	discriminant := halfB*halfB - a*c

	if discriminant > 0 {
		root := math.Sqrt(discriminant)

		temp := (-halfB - root) / a
		if temp > tMin && temp < tMax {
			rec.Rec.T = temp
			rec.Rec.P = r.At(rec.Rec.T)
			rec.Rec.U, rec.Rec.V = getSphereUV(maths.Sub(rec.Rec.P, s.Center).Div(s.Radius))
			outwardNormal := maths.Sub(rec.Rec.P, s.Center).Div(s.Radius)
			rec.Rec.SetFaceNormal(r, outwardNormal)
			rec.Mat = s.Mat
			return true
		}

		temp = (-halfB + root) / a
		if temp > tMin && temp < tMax {
			rec.Rec.T = temp
			rec.Rec.P = r.At(rec.Rec.T)
			rec.Rec.U, rec.Rec.V = getSphereUV(maths.Sub(rec.Rec.P, s.Center).Div(s.Radius))
			outwardNormal := maths.Sub(rec.Rec.P, s.Center).Div(s.Radius)
			rec.Rec.SetFaceNormal(r, outwardNormal)
			rec.Mat = s.Mat
			return true
		}
	}
	return false
}

func (s *Sphere) BoundingBox(t0, t1 float64, outputBox *AABB) bool {
	*outputBox = *NewAABB(maths.Sub(s.Center, maths.NewVec3(s.Radius, s.Radius, s.Radius)),
		maths.Add(s.Center, maths.NewVec3(s.Radius, s.Radius, s.Radius)))
	return true
}

func getSphereUV(p *maths.Vec3) (float64, float64) {
	phi := math.Atan2(p.Z(), p.X())
	theta := math.Asin(p.Y())
	u := 1 - (phi+math.Pi)/(2*math.Pi)
	v := (theta + math.Pi/2) / math.Pi
	return u, v
}
