package moretypes

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8
	for i := 0; i < dy; i++ {
		var x []uint8
		for j := 0; j < dx; j++ {
			x = append(x, 0)
		}
		result = append(result, x)
		fmt.Println(len(result[i]))
	}
	// fmt.Println(result)

	return result
}

func ExerciseSlice() {
	pic.Show(Pic)
}
