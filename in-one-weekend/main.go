package main

import (
	"fmt"
	"os"
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
			r := float64(i) / (imageWidth - 1)
			g := float64(j) / (imageHeight - 1)
			b := 0.25

			ir := int(255 * r)
			ig := int(255 * g)
			ib := int(255 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
	fmt.Fprintf(os.Stderr, "Done\n")
}
