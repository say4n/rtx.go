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

	world := gfx.NewWorld(2)

	world.Hitlist[0] = gfx.Sphere{gfx.Vector3D{0.0, 0.0, -1.0}, 0.5}
	world.Hitlist[1] = gfx.Sphere{gfx.Vector3D{0, -100.5, -1.0}, 100}

	for ny := height - 1; ny >= 0; ny-- {
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
			c := gfx.GetColorFromRay(r, world)

			ir := int(255.99 * c.Red)
			ig := int(255.99 * c.Green)
			ib := int(255.99 * c.Blue)

			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}
}
