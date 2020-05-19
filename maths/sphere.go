package maths

import "math"

type Sphere struct {
	Center *Point3
	Radius float64
}

func NewSphere(center *Point3, radius float64) *Sphere {
	return &Sphere{center, radius}
}

func (s *Sphere) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := Sub(r.Origin(), s.Center)
	a := r.Direction().LenSquared()
	halfB := Dot(oc, r.Direction())
	c := oc.LenSquared() - s.Radius*s.Radius
	discriminant := halfB*halfB - a*c

	if discriminant > 0 {
		root := math.Sqrt(discriminant)

		temp := (-halfB - root) / a
		if temp > tMin && temp < tMax {
			rec.T = temp
			rec.P = r.At(rec.T)
			rec.Normal = Sub(rec.P, s.Center).Div(s.Radius)
			return true
		}

		temp = (-halfB + root) / a
		if temp > tMin && temp < tMax {
			rec.T = temp
			rec.P = r.At(rec.T)
			rec.Normal = Sub(rec.P, s.Center).Div(s.Radius)
			return true
		}
	}
	return false
}
