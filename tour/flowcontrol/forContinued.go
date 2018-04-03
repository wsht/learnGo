package flowcontrol

import "fmt"

func ForContinued() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
