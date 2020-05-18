package main

import (
	"fmt"
	"os"

	"github.com/iCiaran/ray-tracing/maths"
	"github.com/iCiaran/ray-tracing/util/colour"
)

const (
	imageHeight = 256
	imageWidth  = 256
)

func main() {
	fmt.Printf("P3\n%d %d\n255\n", imageHeight, imageWidth)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "Scanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			pixelColour := maths.NewVec3(float64(i)/(imageWidth-1), float64(j)/(imageHeight-1), 0.25)
			colour.WriteColour(os.Stdout, pixelColour)
		}
	}
	fmt.Fprintf(os.Stderr, "Done\n")
}
