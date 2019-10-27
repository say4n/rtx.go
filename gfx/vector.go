package gfx

import "math"

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func NewVector3D(x, y, z float64) *Vector3D {
	var vec Vector3D

	vec.X = x
	vec.Y = y
	vec.Z = z

	return &vec
}

func (v *Vector3D) GetX() float64 {
	return v.X
}

func (v *Vector3D) GetY() float64 {
	return v.Y
}

func (v *Vector3D) GetZ() float64 {
	return v.Z
}

func (v *Vector3D) GetLength() float64 {
	x := v.X
	y := v.Y
	z := v.Z

	return math.Sqrt(x*x + y*y + z*z)
}

func (v *Vector3D) GetSquaredLength() float64 {
	x := v.X
	y := v.Y
	z := v.Z

	return x*x + y*y + z*z
}

func (v *Vector3D) Scale(scalingFactor float64) *Vector3D {
	var vec Vector3D

	vec.X = v.X * scalingFactor
	vec.Y = v.Y * scalingFactor
	vec.Z = v.Z * scalingFactor

	return &vec
}

func (v *Vector3D) Normalize() *Vector3D {
	var vec Vector3D

	scalingFactor := 1.0 / v.GetLength()
	vec = *v.Scale(scalingFactor)

	return &vec
}

func (v *Vector3D) Negate() *Vector3D {
	var vec Vector3D

	vec.X = -v.X
	vec.Y = -v.Y
	vec.Z = -v.Z

	return &vec
}

func (v1 *Vector3D) Add(v2 *Vector3D) *Vector3D {
	var vec Vector3D

	vec.X = v1.X + v2.X
	vec.Y = v1.Y + v2.Y
	vec.Z = v1.Z + v2.Z

	return &vec
}

func (v1 *Vector3D) Subtract(v2 *Vector3D) *Vector3D {
	var vec Vector3D

	vec.X = v1.X - v2.X
	vec.Y = v1.Y - v2.Y
	vec.Z = v1.Z - v2.Z

	return &vec
}

func (v1 *Vector3D) InnerProduct(v2 *Vector3D) *Vector3D {
	var vec Vector3D

	vec.X = v1.X * v2.X
	vec.Y = v1.Y * v2.Y
	vec.Z = v1.Z * v2.Z

	return &vec
}

func (v1 *Vector3D) DotProduct(v2 *Vector3D) float64 {
	var pdt float64
	pdt = v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z

	return pdt
}

func (v1 *Vector3D) CrossProduct(v2 *Vector3D) *Vector3D {
	var vec Vector3D

	vec.X = v1.Y*v2.Z - v1.Z*v1.Y
	vec.Y = v1.Z*v2.X - v1.X*v1.Z
	vec.Z = v1.X*v2.Y - v1.Y*v2.X

	return &vec
}

func (v1 *Vector3D) Divide(v2 *Vector3D) *Vector3D {
	var vec Vector3D

	vec.X = v1.X / v2.X
	vec.Y = v1.Y / v2.Y
	vec.Z = v1.Z / v2.Z

	return &vec
}
