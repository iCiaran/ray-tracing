package maths

type HitRecord struct {
	P      *Point3
	Normal *Vec3
	T      float64
}

func NewHitRecord(p *Point3, normal *Vec3, t float64) *HitRecord {
	return &HitRecord{p, normal, t}
}

type hittable interface {
	Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool
}
