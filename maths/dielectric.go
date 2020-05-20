package maths

import "math"

type Dielectric struct {
	refIdx float64
}

func NewDielectric(refIdx float64) *Dielectric {
	return &Dielectric{refIdx}
}

func (d *Dielectric) Scatter(rIn *Ray, rec *HitRecord, attenuation *Colour, scattered *Ray) bool {
	*attenuation = *NewVec3(1.0, 1.0, 1.0)

	etaiOverEtat := d.refIdx
	if rec.FrontFace {
		etaiOverEtat = 1.0 / etaiOverEtat
	}

	unitDirection := Normalise(rIn.Direction())
	cosTheta := math.Min(Dot(Neg(unitDirection), rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	if etaiOverEtat*sinTheta > 1.0 || Random() < Schlick(cosTheta, etaiOverEtat) {
		reflected := Reflect(unitDirection, rec.Normal)
		*scattered = *NewRay(rec.P, reflected)
		return true
	}

	refracted := Refract(unitDirection, rec.Normal, etaiOverEtat)
	*scattered = *NewRay(rec.P, refracted)
	return true
}
