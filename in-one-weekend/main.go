package main

import (
	"fmt"
	"math"
	"os"

	"github.com/iCiaran/ray-tracing/maths"
	"github.com/iCiaran/ray-tracing/util/colour"
)

const (
	aspectRatio = 16.0 / 9.0
	imageWidth  = 384
	imageHeight = int(imageWidth / aspectRatio)
)

func main() {
	cam := maths.NewCamera(aspectRatio, 2.0, 1.0)

	world := maths.NewHittableList()
	world.Add(maths.NewSphere(maths.NewVec3(0.0, 0.0, -1.0), 0.5))
	world.Add(maths.NewSphere(maths.NewVec3(0.0, -100.5, -1.0), 100))

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "Scanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)
			r := cam.GetRay(u, v)
			pixelColour := rayColour(r, world)
			colour.WriteColour(os.Stdout, pixelColour)
		}
	}
	fmt.Fprintf(os.Stderr, "Done\n")
}

func rayColour(r *maths.Ray, world maths.Hittable) *maths.Colour {
	rec := maths.NewHitRecord()
	if world.Hit(r, 0, math.Inf(1), rec) {
		return maths.Add(rec.Normal, maths.NewVec3(1.0, 1.0, 1.0)).Mul(0.5)
	}

	unitDirection := maths.Normalise(r.Direction())
	t := 0.5 * (unitDirection.Y() + 1.0)
	return maths.Add(maths.Mul(&maths.Colour{1.0, 1.0, 1.0}, (1.0-t)), maths.Mul(&maths.Colour{0.5, 0.7, 1.0}, t))
}
