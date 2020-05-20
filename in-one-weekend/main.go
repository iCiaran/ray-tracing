package main

import (
	"fmt"
	"math"
	"os"

	"github.com/iCiaran/ray-tracing/maths"
)

const (
	aspectRatio     = 16.0 / 9.0
	imageWidth      = 384
	imageHeight     = int(imageWidth / aspectRatio)
	samplesPerPixel = 100
	maxDepth        = 50
)

func main() {
	cam := maths.NewCamera(aspectRatio, 2.0, 1.0)

	world := maths.NewHittableList()
	world.Add(maths.NewSphere(maths.NewVec3(0.0, 0.0, -1.0), 0.5, maths.NewLambertian(maths.NewVec3(0.1, 0.2, 0.5))))
	world.Add(maths.NewSphere(maths.NewVec3(0.0, -100.5, -1.0), 100, maths.NewLambertian(maths.NewVec3(0.8, 0.8, 0.0))))
	world.Add(maths.NewSphere(maths.NewVec3(1.0, 0, -1.0), 0.5, maths.NewMetal(maths.NewVec3(0.8, 0.6, 0.2), 0.0)))
	world.Add(maths.NewSphere(maths.NewVec3(-1.0, 0, -1.0), 0.5, maths.NewDielectric(1.5)))
	world.Add(maths.NewSphere(maths.NewVec3(-1.0, 0, -1.0), -0.45, maths.NewDielectric(1.5)))

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "Scanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			pixelColour := maths.NewVec3(0.0, 0.0, 0.0)
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + maths.Random()) / float64(imageWidth-1)
				v := (float64(j) + maths.Random()) / float64(imageHeight-1)
				r := cam.GetRay(u, v)
				pixelColour.Add(rayColour(r, world, maxDepth))
			}
			maths.WriteColour(os.Stdout, pixelColour, samplesPerPixel)
		}
	}
	fmt.Fprintf(os.Stderr, "Done\n")
}

func rayColour(r *maths.Ray, world maths.Hittable, depth int) *maths.Colour {
	if depth <= 0 {
		return maths.NewVec3(0.0, 0.0, 0.0)
	}

	rec := maths.NewHitRecord()
	if world.Hit(r, 0.001, math.Inf(1), rec) {
		scattered := maths.NewEmptyRay()
		attenuation := maths.NewVec3(0.0, 0.0, 0.0)

		if rec.Mat.Scatter(r, rec, attenuation, scattered) {
			return maths.MulVec(attenuation, rayColour(scattered, world, depth-1))
		}
		return maths.NewVec3(0.0, 0.0, 0.0)
	}

	unitDirection := maths.Normalise(r.Direction())
	t := 0.5 * (unitDirection.Y() + 1.0)
	return maths.Add(maths.Mul(&maths.Colour{1.0, 1.0, 1.0}, (1.0-t)), maths.Mul(&maths.Colour{0.5, 0.7, 1.0}, t))
}
