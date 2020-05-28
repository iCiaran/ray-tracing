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

	textureA := maths.NewTextureSolid(0.2, 0.3, 0.1)
	textureB := maths.NewTextureSolid(0.9, 0.9, 0.9)

	texture := maths.NewTextureChecker(textureA, textureB, 0.1)

	world := maths.NewHittableList()
	world.Add(maths.NewSphere(maths.NewVec3(0.0, -10.0, 0.0), 10.0, maths.NewLambertian(texture)))
	world.Add(maths.NewSphere(maths.NewVec3(0.0, 10.0, 0.0), 10.0, maths.NewLambertian(texture)))

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
