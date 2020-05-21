package maths

import "math"

type Camera struct {
	origin          *Point3
	lowerLeftCorner *Point3
	horizontal      *Vec3
	vertical        *Vec3
}

func NewCamera(lookFrom, lookAt, up *Vec3, vFov, aspectRatio float64) *Camera {
	cam := Camera{}

	theta := DegreesToRadians(vFov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := Sub(lookFrom, lookAt).Normalise()
	u := Cross(up, w).Normalise()
	v := Cross(w, u)

	cam.origin = lookFrom
	cam.horizontal = Mul(u, viewportWidth)
	cam.vertical = Mul(v, viewportHeight)
	cam.lowerLeftCorner = Sub(cam.origin, Div(cam.horizontal, 2.0)).Sub(Div(cam.vertical, 2.0)).Sub(w)

	return &cam
}

func (c *Camera) GetRay(u, v float64) *Ray {
	return NewRay(c.origin, Add(c.lowerLeftCorner, Mul(c.horizontal, u)).Add(Mul(c.vertical, v)).Sub(c.origin))
}
