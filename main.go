package main

import (
	"fmt"
	"os"

	"github.com/say4n/rtx.go/gfx"
)

func main() {
	width := 400
	height := 300

	maxColor := 255.99

	filename := "output.ppm"
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(f, "P3\n%d %d\n%d\n", width, height, int(maxColor))

	for ny := 0; ny < height; ny++ {
		for nx := 0; nx < width; nx++ {
			u := float64(nx) / float64(width)
			v := float64(ny) / float64(height)

			start := gfx.ORIGIN
			end := gfx.BOTTOM_LEFT.Add(
				gfx.HORIZONTAL.Scale(u),
			).Add(
				gfx.VERTICAL.Scale(v),
			)

			r := gfx.Ray{start, end}
			c := gfx.GetColorFromRay(r)

			ir := int(255.99 * c.Red)
			ig := int(255.99 * c.Green)
			ib := int(255.99 * c.Blue)

			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}
}
