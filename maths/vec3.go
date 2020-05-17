package maths

import (
	"fmt"
	"math"
)

type Vec3 [3]float64

type Color = Vec3
type Point3 = Vec3

func NewVec3(e0, e1, e2 float64) *Vec3 {
	return &Vec3{e0, e1, e2}
}

func (v *Vec3) X() float64 {
	return v[0]
}

func (v *Vec3) Y() float64 {
	return v[1]
}

func (v *Vec3) Z() float64 {
	return v[2]
}

func (v *Vec3) Add(a *Vec3) *Vec3 {
	v[0] += a[0]
	v[1] += a[1]
	v[2] += a[2]
	return v
}

func (v *Vec3) Mul(t float64) *Vec3 {
	v[0] *= t
	v[1] *= t
	v[2] *= t
	return v
}

func (v *Vec3) Div(t float64) *Vec3 {
	return v.Mul(1 / t)
}

func (v *Vec3) Neg() *Vec3 {
	return v.Mul(-1)
}

func (v *Vec3) Get(i int) float64 {
	return v[i]
}

func (v *Vec3) Len() float64 {
	return math.Sqrt(v.LenSquared())
}

func (v *Vec3) LenSquared() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v *Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v[0], v[1], v[2])
}
