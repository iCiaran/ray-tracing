package maths

import "github.com/iCiaran/ray-tracing/maths"

type Lambertian struct {
	Albedo texture.Texture
}

func NewLambertian(albedo texture.Texture) *Lambertian {
	return &Lambertian{albedo}
}

func (l *Lambertian) Scatter(rIn *texture.Ray, rec *hittable.HitRecord, attenuation *maths.Colour, scattered *maths.Ray) bool {
	scatterDirection := maths.Add(rec.Normal, maths.RandomUnitVector())
	*scattered = *maths.NewRay(rec.P, scatterDirection)
	*attenuation = *l.Albedo.Value(rec.U, rec.V, rec.P)
	return true
}
