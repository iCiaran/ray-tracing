package maths

type Metal struct {
	Albedo Texture
	fuzz   float64
}

func NewMetal(albedo Texture, fuzz float64) *Metal {
	return &Metal{albedo, Clamp(fuzz, 0.0, 1.0)}
}

func (m *Metal) Scatter(rIn *Ray, rec *HitRecord, attenuation *Colour, scattered *Ray) bool {
	reflected := Reflect(Normalise(rIn.Direction()), rec.Normal)
	*scattered = *NewRay(rec.P, reflected.Add(Mul(RandomInUnitSphere(), m.fuzz)))
	*attenuation = *m.Albedo.Value(rec.U, rec.V, rec.P)

	return Dot(scattered.Direction(), rec.Normal) > 0
}
