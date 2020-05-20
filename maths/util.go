package maths

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}

	return x
}

func Reflect(vec, normal *Vec3) *Vec3 {
	return Sub(vec, Mul(normal, 2*Dot(vec, normal)))
}

func Refract(uv, n *Vec3, etaiOverEtat float64) *Vec3 {
	cosTheta := Dot(Neg(uv), n)
	rOutParallel := Mul(n, cosTheta).Add(uv).Mul(etaiOverEtat)
	rOutPerp := Mul(n, -math.Sqrt(1.0-rOutParallel.LenSquared()))

	return rOutParallel.Add(rOutPerp)
}

func Random() float64 {
	return rand.Float64()
}

func RandomInRange(min, max float64) float64 {
	return min + (max-min)*Random()
}

func RandomPoint() *Vec3 {
	return NewVec3(Random(), Random(), Random())
}

func RandomPointInRange(min, max float64) *Vec3 {
	return NewVec3(RandomInRange(min, max), RandomInRange(min, max), RandomInRange(min, max))
}

func RandomInUnitSphere() *Vec3 {
	for {
		p := RandomPointInRange(-1.0, 1.0)
		if !(p.LenSquared() >= 1.0) {
			return p
		}
	}
}

func RandomUnitVector() *Vec3 {
	a := RandomInRange(0.0, 2*math.Pi)
	z := RandomInRange(-1.0, 1.0)
	r := math.Sqrt(1 - z*z)
	return NewVec3(r*math.Cos(a), r*math.Sin(a), z)
}

func RandomInHemisphere(normal *Vec3) *Vec3 {
	inUnitSphere := RandomInUnitSphere()

	if Dot(inUnitSphere, normal) > 0.0 {
		return inUnitSphere
	}

	return inUnitSphere.Neg()
}

func WriteColour(f *os.File, c *Colour, samplesPerPixel int) {

	scale := 1.0 / float64(samplesPerPixel)

	r := scale * c.R()
	g := scale * c.G()
	b := scale * c.B()

	// Gamma correction for gamma=2.0

	r = math.Sqrt(r)
	g = math.Sqrt(g)
	b = math.Sqrt(b)

	ir := int(256 * Clamp(r, 0.0, 0.999))
	ig := int(256 * Clamp(g, 0.0, 0.999))
	ib := int(256 * Clamp(b, 0.0, 0.999))

	fmt.Printf("%d %d %d\n", ir, ig, ib)
}
