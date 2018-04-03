package moretypes

import (
	"fmt"
	"math"
)

func Function() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}

/**
Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。 函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。

例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上。
*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func FunctionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibinacci() func() int {
	timer := 0
	fn, fn_1 := 1, 1

	return func() int {
		timer++
		if timer < 3 {
			return 1
		} else {
			result := fn + fn_1
			fn_1 = fn
			fn = result
			return result
		}
	}
}

func ExerciseFibonacciClosure() {
	f := fibinacci()
	for i := 0; i <= 10; i++ {
		fmt.Println(f())
	}
}
