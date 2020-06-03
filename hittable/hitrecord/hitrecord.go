package hitrecord

import (
	"github.com/iCiaran/ray-tracing/maths"
)

type HitRecord struct {
	P         *maths.Point3
	Normal    *maths.Vec3
	T         float64
	U         float64
	V         float64
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(r *maths.Ray, outwardNormal *maths.Vec3) {
	h.FrontFace = maths.Dot(r.Direction(), outwardNormal) < 0
	h.Normal = outwardNormal

	if !h.FrontFace {
		h.Normal.Neg()
	}
}

func New() *HitRecord {
	return &HitRecord{maths.NewVec3(0.0, 0.0, 0.0), maths.NewVec3(0.0, 0.0, 0.0), 0.0, 0.0, 0.0, false}
}
