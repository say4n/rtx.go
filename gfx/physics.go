package gfx

var (
	BOTTOM_LEFT = Vector3D{-2.0, -1.0, -1.0}
	HORIZONTAL  = Vector3D{4.0, 0.0, 0.0}
	VERTICAL    = Vector3D{0.0, 2.0, 0.0}
	ORIGIN      = Vector3D{0.0, 0.0, 0.0}
)

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

func GetColorFromRay(r Ray) Color {
	direction := r.GetDirection()
	normalized_direction := direction.Normalize()
	t := 0.5 * (normalized_direction.GetY() + 1.0)

	c1 := Color{1.0, 1.0, 1.0}.Scale(t)
	c2 := Color{0.5, 0.7, 1.0}.Scale(1.0 - t)

	if hitSphere(Vector3D{0.0, 0.0, -1.0}, 0.5, r) {
		return Color{1.0, 0.0, 0.0}
	}

	return c1.Add(c2)
}

func hitSphere(center Vector3D, radius float64, ray Ray) bool {
	origin := ray.GetOrigin()
	direction := ray.GetDirection()
	vector := origin.Subtract(center)

	a := direction.DotProduct(direction)
	b := 2.0 * vector.DotProduct(direction)
	c := vector.DotProduct(vector) - radius*radius

	disc := b*b - 4.0*a*c

	return disc > 0
}
