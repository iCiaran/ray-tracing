package maths

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
	refracted := Refract(unitDirection, rec.Normal, etaiOverEtat)
	*scattered = *NewRay(rec.P, refracted)
	return true
}
