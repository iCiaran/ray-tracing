package maths

type Ray struct {
	origin *Point3
	dir    *Vec3
}

func NewRay(origin, direction *Vec3) *Ray {
	return &Ray{origin, direction}
}

func (r *Ray) Origin() *Point3 {
	return r.origin
}

func (r *Ray) Direction() *Vec3 {
	return r.dir
}

func (r *Ray) At(t float64) *Point3 {
	return Mul(r.dir, t).Add(r.origin)
}
