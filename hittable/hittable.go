package maths

import "github.com/iCiaran/ray-tracing/maths"

type HitRecord struct {
	P         *maths.Point3
	Normal    *maths.Vec3
	T         float64
	U         float64
	V         float64
	FrontFace bool
	Mat       material.Material
}

func (h *HitRecord) SetFaceNormal(r *maths.Ray, outwardNormal *maths.Vec3) {
	h.FrontFace = maths.Dot(r.Direction(), outwardNormal) < 0
	h.Normal = outwardNormal

	if !h.FrontFace {
		h.Normal.Neg()
	}
}

func NewHitRecord() *HitRecord {
	return &HitRecord{maths.NewVec3(0.0, 0.0, 0.0), maths.NewVec3(0.0, 0.0, 0.0), 0.0, 0.0, 0.0, false, nil}
}

type Hittable interface {
	Hit(r *maths.Ray, tMin, tMax float64, rec *HitRecord) bool
	BoundingBox(t0, t1 float64, outputBox *AABB) bool
}
