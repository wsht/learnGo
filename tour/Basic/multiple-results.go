package basic

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func MultipleResult() {
	a, b := swap("hello", "word")
	fmt.Println(a, b)
}
