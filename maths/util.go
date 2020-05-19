package maths

import (
	"fmt"
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

func WriteColour(f *os.File, c *Colour, samplesPerPixel int) {

	scale := 1.0 / float64(samplesPerPixel)

	r := scale * c.R()
	g := scale * c.G()
	b := scale * c.B()

	ir := int(256 * Clamp(r, 0.0, 0.999))
	ig := int(256 * Clamp(g, 0.0, 0.999))
	ib := int(256 * Clamp(b, 0.0, 0.999))

	fmt.Printf("%d %d %d\n", ir, ig, ib)
}
