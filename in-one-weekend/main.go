package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/iCiaran/ray-tracing/maths"
)

const (
	aspectRatio     = 16.0 / 9.0
	imageWidth      = 384
	imageHeight     = int(imageWidth / aspectRatio)
	samplesPerPixel = 100
	maxDepth        = 50
)

var maxBounce = 0

func main() {
	lookFrom := maths.NewVec3(13.0, 2.0, 3.0)
	lookAt := maths.NewVec3(0.0, 0.0, 0.0)
	up := maths.NewVec3(0.0, 1.0, 0.0)
	distToFocus := 10.0
	aperture := 0.1

	cam := maths.NewCamera(lookFrom, lookAt, up, 20.0, aspectRatio, aperture, distToFocus)

	world := randomScene()
	start := time.Now()

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		lineStart := time.Now()
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
		fmt.Fprintf(os.Stderr, "Scanlines remaining: %d -- Last time: %v\n", j, time.Now().Sub(lineStart))
	}
	fmt.Fprintf(os.Stderr, "Done Time: %v -- Max Bounces: %d\n", time.Now().Sub(start), maxBounce)
}

func rayColour(r *maths.Ray, world maths.Hittable, depth int) *maths.Colour {
	if maxDepth-depth > maxBounce {
		maxBounce = maxDepth - depth
	}

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

func randomScene() *maths.HittableList {
	world := maths.NewHittableList()

	groundMaterial := maths.NewLambertian(maths.NewVec3(0.5, 0.5, 0.5))
	world.Add(maths.NewSphere(maths.NewVec3(0.0, -1000.0, 0.0), 1000.0, groundMaterial))
	bigA := maths.NewVec3(0.0, 1.0, 0.0)
	bigB := maths.NewVec3(-4.0, 1.0, 0.0)
	bigC := maths.NewVec3(4.0, 1.0, 0.0)

	matA := maths.NewDielectric(1.5)
	matB := maths.NewLambertian(maths.NewVec3(0.2, 0.0, 0.25))
	matC := maths.NewMetal(maths.NewVec3(0.7, 0.6, 0.5), 0.0)

	world.Add(maths.NewSphere(bigA, 1.0, matA))
	world.Add(maths.NewSphere(bigB, 1.0, matB))
	world.Add(maths.NewSphere(bigC, 1.0, matC))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := maths.Random()

			r := 0.2
			center := maths.NewVec3(float64(a)+0.8*maths.Random(), r, float64(b)+0.8*maths.Random())

			if maths.Sub(center, maths.NewVec3(4.0, r, 0.0)).Len() > 0.9 &&
				maths.Sub(center, maths.NewVec3(0.0, r, 0.0)).Len() > 0.9 &&
				maths.Sub(center, maths.NewVec3(-4.0, r, 0.0)).Len() > 0.9 {

				switch {
				case chooseMat < 0.5:
					albedo := maths.MulVec(maths.RandomPoint(), maths.RandomPoint())
					mat := maths.NewLambertian(albedo)
					world.Add(maths.NewSphere(center, r, mat))
				case chooseMat < 0.85:
					albedo := maths.RandomPointInRange(0.5, 1.0)
					fuzz := maths.RandomInRange(0.0, 0.5)
					mat := maths.NewMetal(albedo, fuzz)
					world.Add(maths.NewSphere(center, r, mat))
				default:
					mat := maths.NewDielectric(1.5)
					world.Add(maths.NewSphere(center, r, mat))
				}
			}
		}
	}

	return world
}
