package gfx

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func NewColor(r, g, b float64) Color {
	var c Color

	c.Red = r
	c.Green = g
	c.Blue = b

	return c
}

func (c Color) Scale(scalingFactor float64) Color {
	var color Color

	color.Red = c.Red * scalingFactor
	color.Green = c.Green * scalingFactor
	color.Blue = c.Blue * scalingFactor

	return color
}

func (c1 Color) Add(c2 Color) Color {
	var color Color

	color.Red = c1.Red + c2.Red
	color.Green = c1.Green + c2.Green
	color.Blue = c1.Blue + c2.Blue

	return color
}

func (c Color) Divide(d float64) Color {
	var color Color

	color.Red = c.Red / d
	color.Green = c.Green / d
	color.Blue = c.Blue / d

	return color
}
