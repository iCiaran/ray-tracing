package maths

import "math"

type Camera struct {
	origin          *Point3
	lowerLeftCorner *Point3
	horizontal      *Vec3
	vertical        *Vec3
}

func NewCamera(aspectRatio, focalLength, vFov float64) *Camera {
	cam := Camera{}

	theta := DegreesToRadians(vFov)
	h := math.Tan(theta / 2.0)

	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	cam.origin = NewVec3(0.0, 0.0, 0.0)
	cam.horizontal = NewVec3(viewportWidth, 0.0, 0.0)
	cam.vertical = NewVec3(0.0, viewportHeight, 0.0)
	cam.lowerLeftCorner = Sub(cam.origin, Div(cam.horizontal, 2.0)).Sub(Div(cam.vertical, 2.0)).Sub(NewVec3(0.0, 0.0, focalLength))

	return &cam
}

func (c *Camera) GetRay(u, v float64) *Ray {
	return NewRay(c.origin, Add(c.lowerLeftCorner, Mul(c.horizontal, u)).Add(Mul(c.vertical, v)).Sub(c.origin))
}
