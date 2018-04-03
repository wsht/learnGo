package flowcontrol

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, float64) {
	z := float64(1)
	// y := float64(1)

	for i := 0; i < 1000; i++ {
		z = z - (z*z-x)/2*z
		// fmt.Println(z)
	}

	return z, math.Sqrt(4)
}

func ExerciseLoopAndFunctions() {
	fmt.Println(Sqrt(4))
}
