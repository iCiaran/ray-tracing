package maths

type Lambertian struct {
	Albedo *Colour
}

func NewLambertian(albedo *Colour) *Lambertian {
	return &Lambertian{albedo}
}

func (l *Lambertian) Scatter(rIn *Ray, rec *HitRecord, attenuation *Colour, scattered *Ray) bool {
	scatterDirection := Add(rec.Normal, RandomUnitVector())
	*scattered = *NewRay(rec.P, scatterDirection)
	*attenuation = *l.Albedo
	return true
}
