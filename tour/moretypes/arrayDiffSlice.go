package moretypes

import (
	"fmt"
)

/**
并且数组返回的是值类型，把一个数组赋予给另一个数组时发生的是值拷贝
而切片是指针类型，拷贝的是指针
*/

func ArrayDiffSlice() {
	var a = [3]int{1, 2, 3}

	b := a
	b[1] = 10
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	c := []int{1, 2, 3}
	d := c
	d[1] = 10

	fmt.Println("c:", c)
	fmt.Println("d:", d)

	fmt.Println(&d)
}
