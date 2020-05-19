package maths

type HittableList struct {
	objects []Hittable
}

func NewHittableList() *HittableList {
	return &HittableList{make([]Hittable, 0)}
}

func (l *HittableList) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool {
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

func (l *HittableList) Add(object Hittable) {
	l.objects = append(l.objects, object)
}

func (l *HittableList) Clear() {
	l.objects = nil
}
