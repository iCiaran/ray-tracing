package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"time"

	"github.com/iCiaran/ray-tracing/maths"
)

var maxBounce = 0

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	aspectRatio := config.Aspect.Width / config.Aspect.Height
	imageWidth := config.ImageWidth
	imageHeight := int(float64(imageWidth) / aspectRatio)
	samplesPerPixel := config.SamplesPerPixel
	maxDepth := config.MaxDepth
	lookFrom := maths.NewVec3(config.Camera.From.X, config.Camera.From.Y, config.Camera.From.Z)
	lookAt := maths.NewVec3(config.Camera.To.X, config.Camera.To.Y, config.Camera.To.Z)
	up := maths.NewVec3(config.Camera.Up.X, config.Camera.Up.Y, config.Camera.Up.Z)
	distToFocus := config.Camera.DistToFocus
	aperture := config.Camera.Aperture
	vFov := config.Camera.VFov

	cam := maths.NewCamera(lookFrom, lookAt, up, vFov, aspectRatio, aperture, distToFocus)

	world := randomScene()

	start := time.Now()

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		lineStart := time.Now()
		for i := 0; i < imageWidth; i++ {
			pixelColour := conPixelColour(i, j, imageWidth, imageHeight, samplesPerPixel, maxDepth, world, cam)
			maths.WriteColour(os.Stdout, pixelColour, samplesPerPixel)
		}
		fmt.Fprintf(os.Stderr, "Scanlines remaining: %d -- Last time: %v\n", j, time.Now().Sub(lineStart))
	}
	fmt.Fprintf(os.Stderr, "Done Time: %v -- Max Bounces: %d\n", time.Now().Sub(start), maxBounce)
}

func rayColour(r *maths.Ray, world maths.Hittable, maxDepth, depth int) *maths.Colour {
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
			return maths.MulVec(attenuation, rayColour(scattered, world, maxDepth, depth-1))
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
				case chooseMat < 0.3:
					albedo := maths.MulVec(maths.RandomPoint(), maths.RandomPoint())
					mat := maths.NewLambertian(albedo)
					world.Add(maths.NewSphere(center, r, mat))
				case chooseMat < 0.6:
					albedo := maths.RandomPointInRange(0.5, 1.0)
					fuzz := maths.RandomInRange(0.0, 0.5)
					mat := maths.NewMetal(albedo, fuzz)
					world.Add(maths.NewSphere(center, r, mat))
				case chooseMat < 0.7:
					mat := maths.NewDielectric(1.5)
					world.Add(maths.NewSphere(center, r, mat))
				case chooseMat < 0.8:
					albedo := maths.MulVec(maths.RandomPoint(), maths.RandomPoint())
					mat := maths.NewLambertian(albedo)
					addCube(world, r, center, mat)
				case chooseMat < 0.9:
					albedo := maths.RandomPointInRange(0.5, 1.0)
					fuzz := maths.RandomInRange(0.0, 0.5)
					mat := maths.NewMetal(albedo, fuzz)
					addCube(world, r, center, mat)
				case chooseMat < 1.0:
					mat := maths.NewDielectric(1.5)
					addCube(world, r, center, mat)
				}
			}
		}
	}

	return world
}

func addCube(world *maths.HittableList, size float64, center *maths.Point3, mat maths.Material) {
	v1 := maths.NewVec3(-size, size, size).Add(center)   // left front top
	v2 := maths.NewVec3(size, size, size).Add(center)    // right front top
	v3 := maths.NewVec3(-size, -size, size).Add(center)  // left front bottom
	v4 := maths.NewVec3(size, -size, size).Add(center)   // right front bottom
	v5 := maths.NewVec3(-size, size, -size).Add(center)  // left back top
	v6 := maths.NewVec3(size, size, -size).Add(center)   // right back top
	v7 := maths.NewVec3(-size, -size, -size).Add(center) // left back bottom
	v8 := maths.NewVec3(size, -size, -size).Add(center)  // right back bottom
	nFront := maths.NewVec3(0.0, 0.0, 1.0)
	nBack := maths.NewVec3(0.0, 0.0, -1.0)
	nLeft := maths.NewVec3(-1.0, 0.0, 0.0)
	nRight := maths.NewVec3(1.0, 0.0, 0.0)
	nUp := maths.NewVec3(0.0, 1.0, 0.0)
	nDown := maths.NewVec3(0.0, -1.0, 0.0)

	// Front v3 v4 v2 / v2 v1 v3
	world.Add(maths.NewTriangle(v3, v4, v2, nFront, nFront, nFront, mat, true))
	world.Add(maths.NewTriangle(v2, v1, v3, nFront, nFront, nFront, mat, true))

	// Back v8 v7 v5 / v5 v6 v8
	world.Add(maths.NewTriangle(v8, v7, v5, nBack, nBack, nBack, mat, true))
	world.Add(maths.NewTriangle(v5, v6, v8, nBack, nBack, nBack, mat, true))

	// Right v4 v8 v6 / v6 v2 v4
	world.Add(maths.NewTriangle(v4, v8, v6, nRight, nRight, nRight, mat, true))
	world.Add(maths.NewTriangle(v6, v2, v4, nRight, nRight, nRight, mat, true))

	// Left v7 v3 v1 / v1 v5 v7
	world.Add(maths.NewTriangle(v7, v3, v1, nLeft, nLeft, nLeft, mat, true))
	world.Add(maths.NewTriangle(v1, v5, v7, nLeft, nLeft, nLeft, mat, true))

	// Top v1 v2 v6 / v6 v5 v1
	world.Add(maths.NewTriangle(v1, v2, v6, nUp, nUp, nUp, mat, true))
	world.Add(maths.NewTriangle(v6, v5, v1, nUp, nUp, nUp, mat, true))

	// Bottom v4 v3 v7 / v7 v8 v4
	world.Add(maths.NewTriangle(v4, v3, v7, nDown, nDown, nDown, mat, true))
	world.Add(maths.NewTriangle(v7, v8, v4, nDown, nDown, nDown, mat, true))
}

func pixelColour(i, j, imageWidth, imageHeight, samplesPerPixel, maxDepth int, world *maths.HittableList, cam *maths.Camera) *maths.Colour {
	pixelColour := maths.NewVec3(0.0, 0.0, 0.0)
	for s := 0; s < samplesPerPixel; s++ {
		u := (float64(i) + maths.Random()) / float64(imageWidth-1)
		v := (float64(j) + maths.Random()) / float64(imageHeight-1)
		r := cam.GetRay(u, v)
		pixelColour.Add(rayColour(r, world, maxDepth, maxDepth))
	}
	return pixelColour
}

func conPixelColour(i, j, imageWidth, imageHeight, samplesPerPixel, maxDepth int, world *maths.HittableList, cam *maths.Camera) *maths.Colour {
	pixelColour := maths.NewVec3(0.0, 0.0, 0.0)
	c := make(chan *maths.Vec3, samplesPerPixel)
	var wg sync.WaitGroup

	for s := 0; s < samplesPerPixel; s++ {
		wg.Add(1)
		u := (float64(i) + maths.Random()) / float64(imageWidth-1)
		v := (float64(j) + maths.Random()) / float64(imageHeight-1)
		r := cam.GetRay(u, v)
		go conRayColour(r, world, maxDepth, maxDepth, c, &wg)
	}
	wg.Wait()

	for s := 0; s < samplesPerPixel; s++ {
		pixelColour.Add(<-c)
	}

	return pixelColour
}

func conRayColour(r *maths.Ray, world maths.Hittable, maxDepth, depth int, c chan *maths.Vec3, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- rayColour(r, world, maxDepth, depth)
}
