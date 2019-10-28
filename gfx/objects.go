package gfx

import "math"

type World struct {
	Size    int
	Hitlist []hittable
}

type hittable interface {
	hit(Ray, float64, float64, *hitRecord) bool
}

func NewWorld(size int) *World {
	w := World{}
	w.Size = size
	w.Hitlist = make([]hittable, 2)

	return &w
}

func (w *World) hit(ray Ray, t_min, t_max float64, hit *hitRecord) bool {
	var tempRec hitRecord
	hitAnything := false
	closestSoFar := t_max

	for _, h := range w.Hitlist {
		if h.hit(ray, t_min, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			*hit = tempRec
		}
	}

	return hitAnything
}

type Sphere struct {
	Center Vector3D
	Radius float64
}

func (s Sphere) hit(ray Ray, t_min, t_max float64, hit *hitRecord) bool {
	origin := ray.GetOrigin()
	direction := ray.GetDirection()
	vector := origin.Subtract(s.Center)

	a := direction.DotProduct(direction)
	b := 2.0 * vector.DotProduct(direction)
	c := vector.DotProduct(vector) - s.Radius*s.Radius

	disc := b*b - 4.0*a*c

	if disc > 0.0 {
		t := (-b - math.Sqrt(disc)) / (2.0 * a)
		if t < t_max && t > t_min {
			hit.t = t
			hit.point = ray.GetPointAtParameter(t)
			vector := hit.point.Subtract(s.Center)
			hit.normal = vector.Normalize()

			return true
		}

		t = (-b + math.Sqrt(disc)) / (2.0 * a)
		if t < t_max && t > t_min {
			hit.t = t
			hit.point = ray.GetPointAtParameter(t)
			vector := hit.point.Subtract(s.Center)
			hit.normal = vector.Normalize()

			return true
		}
	}

	return false
}
