package gfx

type Ray struct {
	A Vector3D
	B Vector3D
}

func (r *Ray) GetOrigin() Vector3D {
	return r.A
}

func (r *Ray) GetDirection() Vector3D {
	return r.B
}

func (r *Ray) GetPointAtParameter(t float64) Vector3D {
	return r.B.Scale(t)
}

type hitRecord struct {
	t      float64
	point  Vector3D
	normal Vector3D
}

func GetColorFromRay(r Ray, world *World) Color {
	var hit hitRecord

	if world.hit(r, 0.0, FLOATMAX, &hit) {
		normal := hit.normal

		r := 0.5 * (normal.GetX() + 1.0)
		g := 0.5 * (normal.GetY() + 1.0)
		b := 0.5 * (normal.GetZ() + 1.0)

		return Color{r, g, b}
	} else {
		direction := r.GetDirection()
		normalized_direction := direction.Normalize()
		t := 0.5 * (normalized_direction.GetY() + 1.0)

		c1 := Color{1.0, 1.0, 1.0}.Scale(1.0 - t)
		c2 := Color{0.5, 0.7, 1.0}.Scale(t)

		return c1.Add(c2)
	}
}
