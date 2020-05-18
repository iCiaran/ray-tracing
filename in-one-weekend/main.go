package main

import (
	"fmt"
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
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := maths.NewVec3(0.0, 0.0, 0.0)
	horizontal := maths.NewVec3(viewportWidth, 0.0, 0.0)
	vertical := maths.NewVec3(0.0, viewportHeight, 0.0)
	lowerLeftCorner := maths.Sub(origin, maths.Div(horizontal, 2.0)).Sub(maths.Div(vertical, 2.0)).Sub(maths.NewVec3(0.0, 0.0, focalLength))

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "Scanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)
			r := maths.NewRay(origin, maths.Add(lowerLeftCorner, maths.Mul(horizontal, u)).Add(maths.Mul(vertical, v)).Sub(origin))
			pixelColour := rayColour(r)
			colour.WriteColour(os.Stdout, pixelColour)
		}
	}
	fmt.Fprintf(os.Stderr, "Done\n")
}

func rayColour(r *maths.Ray) *maths.Colour {
	if hitSphere(maths.NewVec3(0.0, 0.0, -1.0), 0.5, r) {
		return maths.NewVec3(1.0, 0.0, 0.0)
	}
	unitDirection := maths.Normalise(r.Direction())
	t := 0.5 * (unitDirection.Y() + 1.0)
	return maths.Add(maths.Mul(&maths.Colour{1.0, 1.0, 1.0}, (1.0-t)), maths.Mul(&maths.Colour{0.5, 0.7, 1.0}, t))
}

func hitSphere(center *maths.Point3, radius float64, r *maths.Ray) bool {
	oc := maths.Sub(r.Origin(), center)
	a := maths.Dot(r.Direction(), r.Direction())
	b := 2.0 * maths.Dot(oc, r.Direction())
	c := maths.Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}
