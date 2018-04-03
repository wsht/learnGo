package flowcontrol

import (
	"fmt"
	"math"
)

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
		if x := true; x {
			fmt.Printf("%g >= %g\n", v, lim)
		}
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func IfAndElse() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
		pow(3, 3, 10),
	)
}
