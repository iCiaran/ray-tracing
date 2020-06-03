package hittable

import (
	"github.com/iCiaran/ray-tracing/hittable/hitmatrecord"
	"github.com/iCiaran/ray-tracing/maths"
)

type List struct {
	objects []Hittable
}

func NewList() *List {
	return &List{make([]Hittable, 0)}
}

func (l *List) Hit(r *maths.Ray, tMin, tMax float64, rec *hitmatrecord.HitMatRecord) bool {
	tempRec := hitmatrecord.New()
	hitAnything := false
	closest := tMax

	for _, o := range l.objects {
		if o.Hit(r, tMin, closest, tempRec) {
			hitAnything = true
			closest = tempRec.Rec.T
			*rec = *tempRec
		}
	}

	return hitAnything
}

func (l *List) BoundingBox(t0, t1 float64, outputBox *AABB) bool {
	if len(l.objects) == 0 {
		return false
	}

	tempBox := NewAABB(maths.NewVec3(0.0, 0.0, 0.0), maths.NewVec3(0.0, 0.0, 0.0))
	firstBox := true

	for _, o := range l.objects {
		if !o.BoundingBox(t0, t1, tempBox) {
			return false
		}

		if firstBox {
			*outputBox = *tempBox
		} else {
			*outputBox = *surroundingBox(outputBox, tempBox)
			firstBox = false
		}

	}
	return true
}

func (l *List) Add(object Hittable) {
	l.objects = append(l.objects, object)
}

func (l *List) Clear() {
	l.objects = nil
}
