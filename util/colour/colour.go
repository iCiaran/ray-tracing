package colour

import (
	"fmt"
	"os"

	"github.com/iCiaran/ray-tracing/maths"
)

func WriteColour(f *os.File, c *maths.Colour) {
	ir := int(255 * c.R())
	ig := int(255 * c.G())
	ib := int(255 * c.B())

	fmt.Printf("%d %d %d\n", ir, ig, ib)
}
