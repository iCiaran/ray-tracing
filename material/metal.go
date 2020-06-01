package maths

import "github.com/iCiaran/ray-tracing/maths"

type Metal struct {
	Albedo texture.Texture
	fuzz   float64
}

func NewMetal(albedo texture.Texture, fuzz float64) *Metal {
	return &Metal{albedo, maths.Clamp(fuzz, 0.0, 1.0)}
}

func (m *Metal) Scatter(rIn *maths.Ray, rec *hittable.HitRecord, attenuation *maths.Colour, scattered *maths.Ray) bool {
	reflected := maths.Reflect(maths.Normalise(rIn.Direction()), rec.Normal)
	*scattered = *maths.NewRay(rec.P, reflected.Add(maths.Mul(maths.RandomInUnitSphere(), m.fuzz)))
	*attenuation = *m.Albedo.Value(rec.U, rec.V, rec.P)

	return maths.Dot(scattered.Direction(), rec.Normal) > 0
}
