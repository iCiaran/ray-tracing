package maths

import "math"

type Sphere struct {
	Center *Point3
	Radius float64
	Mat    Material
}

func NewSphere(center *Point3, radius float64, mat Material) *Sphere {
	return &Sphere{center, radius, mat}
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
			rec.U, rec.V = getSphereUV(Sub(rec.P, s.Center).Div(s.Radius))
			rec.P = r.At(rec.T)
			outwardNormal := Sub(rec.P, s.Center).Div(s.Radius)
			rec.SetFaceNormal(r, outwardNormal)
			rec.Mat = s.Mat
			return true
		}

		temp = (-halfB + root) / a
		if temp > tMin && temp < tMax {
			rec.T = temp
			rec.U, rec.V = getSphereUV(Sub(rec.P, s.Center).Div(s.Radius))
			rec.P = r.At(rec.T)
			outwardNormal := Sub(rec.P, s.Center).Div(s.Radius)
			rec.SetFaceNormal(r, outwardNormal)
			rec.Mat = s.Mat
			return true
		}
	}
	return false
}

func (s *Sphere) BoundingBox(t0, t1 float64, outputBox *AABB) bool {
	*outputBox = *NewAABB(Sub(s.Center, NewVec3(s.Radius, s.Radius, s.Radius)),
		Add(s.Center, NewVec3(s.Radius, s.Radius, s.Radius)))
	return true
}

func getSphereUV(p *Vec3) (float64, float64) {
	phi := math.Atan2(p.Z(), p.X())
	theta := math.Asin(p.Y())
	u := 1 - (phi+math.Pi)/(2*math.Pi)
	v := (theta + math.Pi/2) / math.Pi
	return u, v
}
