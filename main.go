package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/say4n/rtx.go/gfx"
)

func main() {
	width := 400
	height := 300
	numSamplesAA := 100

	r := rand.New(rand.NewSource(42))

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
			color := gfx.Color{0.0, 0.0, 0.0}

			for s := 0; s < numSamplesAA; s++ {
				u := (float64(nx) + r.Float64()) / float64(width)
				v := (float64(ny) + r.Float64()) / float64(height)

				ray := gfx.GetRay(u, v)
				c := gfx.GetColorFromRay(ray, world)

				color = color.Add(c)
			}

			color = color.Divide(float64(numSamplesAA))

			// gamma correction
			color.Red = math.Sqrt(color.Red)
			color.Green = math.Sqrt(color.Green)
			color.Blue = math.Sqrt(color.Blue)

			ir := int(255.99 * color.Red)
			ig := int(255.99 * color.Green)
			ib := int(255.99 * color.Blue)

			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}
}
