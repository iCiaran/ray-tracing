package material

import (
	"math"

	"github.com/iCiaran/ray-tracing/hittable/hitrecord"
	"github.com/iCiaran/ray-tracing/maths"
)

type Dielectric struct {
	refIdx float64
}

func NewDielectric(refIdx float64) *Dielectric {
	return &Dielectric{refIdx}
}

func (d *Dielectric) Scatter(rIn *maths.Ray, rec *hitrecord.HitRecord, attenuation *maths.Colour, scattered *maths.Ray) bool {
	*attenuation = *maths.NewVec3(1.0, 1.0, 1.0)

	etaiOverEtat := d.refIdx
	if rec.FrontFace {
		etaiOverEtat = 1.0 / etaiOverEtat
	}

	unitDirection := maths.Normalise(rIn.Direction())
	cosTheta := math.Min(maths.Dot(maths.Neg(unitDirection), rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	if etaiOverEtat*sinTheta > 1.0 || maths.Random() < maths.Schlick(cosTheta, etaiOverEtat) {
		reflected := maths.Reflect(unitDirection, rec.Normal)
		*scattered = *maths.NewRay(rec.P, reflected)
		return true
	}

	refracted := maths.Refract(unitDirection, rec.Normal, etaiOverEtat)
	*scattered = *maths.NewRay(rec.P, refracted)
	return true
}
