package texture

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/iCiaran/ray-tracing/maths"
)

type Image struct {
	I      image.Image
	width  int
	height int
}

func NewImage(filepath string) *Image {
	reader, err := os.Open(filepath)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	defer reader.Close()

	config, format, err := image.DecodeConfig(reader)
	fmt.Fprintf(os.Stderr, "%v\n", format)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	reader.Seek(0, 0)
	m, _, err := image.Decode(reader)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	fmt.Fprintf(os.Stderr, "%v\n", config)

	return &Image{m, config.Width, config.Height}
}

func (ti *Image) Value(u, v float64, p *maths.Vec3) *maths.Colour {
	if ti.I == nil {
		return maths.NewVec3(0.0, 1.0, 1.0)
	}

	u = maths.Clamp(u, 0.0, 1.0)
	v = 1.0 - maths.Clamp(v, 0.0, 1.0)

	i := int(float64(ti.width) * u)
	j := int(float64(ti.height) * v)
	r, g, b, _ := ti.I.At(i, j).RGBA()

	scale := 1.0 / 65535.0

	return maths.NewVec3(float64(r)*scale, float64(g)*scale, float64(b)*scale)
}
