package maths

type Lambertian struct {
	Albedo Texture
}

func NewLambertian(albedo Texture) *Lambertian {
	return &Lambertian{albedo}
}

func (l *Lambertian) Scatter(rIn *Ray, rec *HitRecord, attenuation *Colour, scattered *Ray) bool {
	scatterDirection := Add(rec.Normal, RandomUnitVector())
	*scattered = *NewRay(rec.P, scatterDirection)
	*attenuation = *l.Albedo.Value(rec.U, rec.V, rec.P)
	return true
}
