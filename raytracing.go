package main

import (
	"fmt"
	colorpoint "src/color"
	geometry "src/geometrymath"
	rayobj "src/ray"
)

type color geometry.Vec3
type vec3 geometry.Vec3
type ray rayobj.Ray

func rayColor(r ray) color {
	unitDir := r.Direction.UnitVector()
	t := 0.5 * (unitDir.Y() + 1.0)
	tempVec := geometry.NewFromArr([3]float64{1.0, 1.0, 1.0})
	tempVec2 := geometry.NewFromArr([3]float64{0.5, 0.7, 1.0})
	return color(geometry.AddV(tempVec.MulScalar(1.0-t), tempVec2.MulScalar(t)))

}

func main() {

	//image

	const ratio = 16.0 / 9.0
	const width = 400
	const height = int(width / ratio)

	//camera

	viewportH := 2.0
	viewportW := ratio * viewportH
	focalLength := 1.0
	origin := geometry.NewVec3()
	horizontal := geometry.NewFromArr([3]float64{viewportW, 0.0, 0.0})
	vertical := geometry.NewFromArr([3]float64{0, viewportH, 0})
	//lower left corner
	stub := geometry.SubV(horizontal.DivScalar(2), vertical.DivScalar(2))
	stub = geometry.SubV(stub, geometry.NewFromArr([3]float64{0, 0, focalLength}))
	llc := geometry.SubV(origin, stub)

	fmt.Printf("P3\n%d %d\n255\n", width, height)
	for j := height - 1; j >= 0; j-- {
		for i := 0; i <= width; i++ {
			u := float64(i) / float64((width - 1))
			v := float64(j) / float64((height - 1))
			dir := geometry.SubV(geometry.AddV(llc, geometry.AddV(horizontal.MulScalar(u), vertical.MulScalar(v))), origin)
			r := rayobj.NewRay(origin, dir)
			pixelcolor := rayColor(ray(r))
			colorpoint.WriteColor(geometry.Vec3(pixelcolor))
		}
	}

}
