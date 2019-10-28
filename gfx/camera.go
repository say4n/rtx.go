package gfx

import (
	"math"
	"math/rand"
)

var (
	BOTTOM_LEFT = Vector3D{-2.0, -1.0, -1.0}
	HORIZONTAL  = Vector3D{4.0, 0.0, 0.0}
	VERTICAL    = Vector3D{0.0, 2.0, 0.0}
	ORIGIN      = Vector3D{0.0, 0.0, 0.0}
	FLOATMAX    = math.MaxFloat64
)

func GetRay(u, v float64) Ray {
	start := ORIGIN
	end := BOTTOM_LEFT.Add(HORIZONTAL.Scale(u)).Add(VERTICAL.Scale(v))

	return Ray{start, end}
}

func RandomPointInUnitSphere() Vector3D {
	var point Vector3D

	for {
		point.X = 2 * rand.Float64()
		point.Y = 2 * rand.Float64()
		point.Z = 2 * rand.Float64()

		point = point.Subtract(Vector3D{1.0, 1.0, 1.0})

		if point.GetSquaredLength() < 1.0 {
			break
		}
	}

	return point
}
