package moretypes

import (
	"fmt"
)

/**
类型 [n]T 是一个有 n 个类型为 T 的值的数组。

表达式

var a [10]int
定义变量 a 是一个有十个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是一个制约，但是请不要担心； Go 提供了更加便利的方式来使用数组。
*/

func Array() {
	var a [2]string
	fmt.Println(a)
	fmt.Println(len(a))
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(len(a))

}

func ArrayInit() {
	arr := [5]int{3: 1, 4: 2}

	fmt.Println(arr)
}
