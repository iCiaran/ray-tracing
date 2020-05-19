package maths

type HitRecord struct {
	P         *Point3
	Normal    *Vec3
	T         float64
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(r *Ray, outwardNormal *Vec3) {
	h.FrontFace = Dot(r.Direction(), outwardNormal) < 0
	h.Normal = outwardNormal

	if !h.FrontFace {
		h.Normal.Neg()
	}
}

func NewHitRecord() *HitRecord {
	return &HitRecord{}
}

type hittable interface {
	Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool
}
