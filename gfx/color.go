package gfx

type Color struct {
	Red   int
	Green int
	Blue  int
}

func NewColor(r, g, b int) *Color {
	var c Color

	c.Red = r
	c.Green = g
	c.Blue = b

	return &c
}
