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
	return *r.B.Scale(t)
}
