package maths

type Metal struct {
	Albedo *Colour
}

func NewMetal(albedo *Colour) *Metal {
	return &Metal{albedo}
}

func (m *Metal) Scatter(rIn *Ray, rec *HitRecord, attenuation *Colour, scattered *Ray) bool {
	reflected := Reflect(Normalise(rIn.Direction()), rec.Normal)
	*scattered = *NewRay(rec.P, reflected)
	*attenuation = *m.Albedo

	return Dot(scattered.Direction(), rec.Normal) > 0
}
