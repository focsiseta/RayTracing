package rayobj

import geometry "src/geometrymath"

type Ray struct {
	Origin    geometry.Vec3
	Direction geometry.Vec3
}

func NewRay(origin, direction geometry.Vec3) Ray {
	return Ray{Origin: origin, Direction: direction}
}

func (ray Ray) at(scalar float64) geometry.Vec3 {
	return geometry.AddV(ray.Origin, ray.Direction.MulScalar(scalar))
}
