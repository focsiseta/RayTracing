package colorpoint

import (
	"fmt"
	geo "src/geometrymath"
)

func WriteColor(color geo.Vec3) {

	fmt.Printf("%d %d %d\n", int(color.X()*255.99), int(color.Y()*255.99), int(color.Z()*255.99))

}
