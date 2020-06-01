package maths

import "github.com/iCiaran/ray-tracing/maths"

type HittableList struct {
	objects []Hittable
}

func NewHittableList() *HittableList {
	return &HittableList{make([]Hittable, 0)}
}

func (l *HittableList) Hit(r *maths.Ray, tMin, tMax float64, rec *HitRecord) bool {
	tempRec := NewHitRecord()
	hitAnything := false
	closest := tMax

	for _, o := range l.objects {
		if o.Hit(r, tMin, closest, tempRec) {
			hitAnything = true
			closest = tempRec.T
			*rec = *tempRec
		}
	}

	return hitAnything
}

func (l *HittableList) BoundingBox(t0, t1 float64, outputBox *AABB) bool {
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

func (l *HittableList) Add(object Hittable) {
	l.objects = append(l.objects, object)
}

func (l *HittableList) Clear() {
	l.objects = nil
}
