package maths

import (
	"fmt"
	"math"
)

type Vec3 [3]float64

type Colour = Vec3
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

func (v *Vec3) R() float64 {
	return v[0]
}

func (v *Vec3) G() float64 {
	return v[1]
}

func (v *Vec3) B() float64 {
	return v[2]
}

func (v *Vec3) Add(a *Vec3) *Vec3 {
	v[0] += a[0]
	v[1] += a[1]
	v[2] += a[2]
	return v
}

func (v *Vec3) Sub(a *Vec3) *Vec3 {
	v[0] -= a[0]
	v[1] -= a[1]
	v[2] -= a[2]
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

func Add(a, b *Vec3) *Vec3 {
	return NewVec3(a[0]+b[0], a[1]+b[1], a[2]+b[2])
}

func Sub(a, b *Vec3) *Vec3 {
	return NewVec3(a[0]-b[0], a[1]-b[1], a[2]-b[2])
}

func Mul(v *Vec3, t float64) *Vec3 {
	return NewVec3(t*v[0], t*v[1], t*v[2])
}

func Div(v *Vec3, t float64) *Vec3 {
	return Mul(v, 1/t)
}

func Dot(u, v *Vec3) float64 {
	return u[0]*v[0] + u[1]*v[1] + u[2]*v[2]
}

func Cross(u, v *Vec3) *Vec3 {
	return NewVec3(
		u[1]*v[2]-u[2]*v[1],
		u[2]*v[0]-u[0]*v[2],
		u[0]*v[1]-u[1]*v[0])
}

func Normalise(v *Vec3) *Vec3 {
	return Div(v, v.Len())
}
