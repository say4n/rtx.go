package main

import (
	"fmt"
	"os"

	"github.com/say4n/rtx.go/gfx"
)

func main() {
	height := 400
	width := 300

	maxColor := 255.99

	filename := "output.ppm"
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(f, "P3\n%d %d\n%d\n", height, width, int(maxColor))

	for ny := 0; ny < height; ny++ {
		for nx := 0; nx < height; nx++ {
			r := float64(nx) / float64(height)
			g := float64(ny) / float64(width)
			b := 0.5

			c := gfx.NewColor(int(maxColor*r), int(maxColor*g), int(maxColor*b))

			fmt.Fprintf(f, "%d %d %d\n", c.Red, c.Green, c.Blue)
		}
	}
}
