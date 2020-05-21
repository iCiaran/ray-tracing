package maths

import "math"

type Camera struct {
	origin          *Point3
	lowerLeftCorner *Point3
	horizontal      *Vec3
	vertical        *Vec3
	w               *Vec3
	u               *Vec3
	v               *Vec3
	lensRadius      float64
}

func NewCamera(lookFrom, lookAt, up *Vec3, vFov, aspectRatio, aperture, focusDist float64) *Camera {
	cam := Camera{}

	theta := DegreesToRadians(vFov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	cam.w = Sub(lookFrom, lookAt).Normalise()
	cam.u = Cross(up, cam.w).Normalise()
	cam.v = Cross(cam.w, cam.u)

	cam.origin = lookFrom
	cam.horizontal = Mul(cam.u, viewportWidth).Mul(focusDist)
	cam.vertical = Mul(cam.v, viewportHeight).Mul(focusDist)
	cam.lowerLeftCorner = Sub(cam.origin, Div(cam.horizontal, 2.0)).Sub(Div(cam.vertical, 2.0)).Sub(Mul(cam.w, focusDist))

	cam.lensRadius = aperture / 2.0

	return &cam
}

func (c *Camera) GetRay(s, t float64) *Ray {
	rd := Mul(RandomInUnitDisk(), c.lensRadius)
	offset := Add(Mul(c.u, rd.X()), Mul(c.v, rd.Y()))
	return NewRay(
		Add(c.origin, offset),
		Add(c.lowerLeftCorner, Mul(c.horizontal, s)).Add(Mul(c.vertical, t)).Sub(c.origin).Sub(offset))
}
