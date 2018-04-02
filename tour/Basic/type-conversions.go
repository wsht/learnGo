package basic

import "math"
import "fmt"

func TypeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)

	fmt.Println(x, y, z)
}
