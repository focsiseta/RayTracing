package geometry

import "math"

type Vec3 struct {
	Vec [3]float64
}

func NewVec3() Vec3 {
	return Vec3{[3]float64{0, 0, 0}}
}
func NewFromArr(vector [3]float64) Vec3 {
	return Vec3{vector}
}

//XYZ
func (v Vec3) X() float64 { return v.Vec[0] }
func (v Vec3) Y() float64 { return v.Vec[1] }
func (v Vec3) Z() float64 { return v.Vec[2] }

//Math op
func (v Vec3) Negate() Vec3 { return Vec3{[3]float64{-v.Vec[0], -v.Vec[1], -v.Vec[2]}} }
func AddV(v1, v2 Vec3) Vec3 {
	return Vec3{
		[3]float64{v1.Vec[0] + v2.Vec[0], v1.Vec[1] + v2.Vec[1], v1.Vec[2] + v2.Vec[2]}}
}

func SubV(v1, v2 Vec3) Vec3 {
	return Vec3{
		[3]float64{v1.Vec[0] - v2.Vec[0], v1.Vec[1] - v2.Vec[1], v1.Vec[2] - v2.Vec[2]}}
}

func (v Vec3) MulScalar(scalar float64) Vec3 {
	return Vec3{[3]float64{scalar * v.Vec[0], scalar * v.Vec[1], scalar * v.Vec[2]}}
}

func (v Vec3) DivScalar(scalar float64) Vec3 {
	return v.MulScalar(1 / scalar)
}

func DotP(v1, v2 Vec3) float64 {

	return v1.Vec[0]*v2.Vec[0] + v1.Vec[1]*v2.Vec[1] + v1.Vec[2]*v2.Vec[2]
}

func (v Vec3) Length() float64 {
	return math.Sqrt(DotP(v, v))
}

func (v Vec3) LengthSquared() float64 {
	return DotP(v, v)
}

func Cross(v1, v2 Vec3) Vec3 {
	return Vec3{
		[3]float64{
			v1.Vec[1]*v2.Vec[2] - v1.Vec[2]*v2.Vec[1],
			v1.Vec[2]*v2.Vec[0] - v1.Vec[0]*v2.Vec[2],
			v1.Vec[0]*v2.Vec[1] - v1.Vec[1]*v2.Vec[0]}}
}

func (v Vec3) UnitVector() Vec3 {
	return v.DivScalar(v.Length())
}
