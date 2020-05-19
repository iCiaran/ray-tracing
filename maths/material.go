package maths

type Material interface {
	Scatter(rIn *Ray, rec *HitRecord, attenuation *Colour, scattered *Ray) bool
}
