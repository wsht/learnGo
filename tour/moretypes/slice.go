package moretypes

import (
	"fmt"
)

/**
一个 slice 会指向一个序列的值，并且包含了长度信息。
[]T 是一个元素类型为 T 的 slice

slice 可以看作一个可以大小可以改变的 数组类型

并且数组返回的是值类型，把一个数组赋予给另一个数组时发生的是值拷贝
而切片是指针类型，拷贝的是指针
*/

func Slice() {
	p := []int{2, 3, 4, 5, 6, 7}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

/**
slice 可以重新切片，创建一个新的 slice 值指向相同的数组。

表达式

s[lo:hi]
表示从 lo 到 hi-1 的 slice 元素，含两端。因此

s[lo:lo]
是空的，而

s[lo:lo+1]
有一个元素。
*/
func SlicingSlice() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])

	fmt.Println("p[3:3] ==", p[3:3])

	fmt.Println("p[3:4] ==", p[3:4])

	p2 := p[3:4]
	p2[0] = 1

	fmt.Println("p ==", p)
	fmt.Println("p2 ==", p2)
}

//make slice
/**
slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组：

a := make([]int, 5)  // len(a)=5
为了指定容量，可传递第三个参数到 `make`：

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/

func MakingSlices() {
	a := make([]int, 5)
	printSlice("a", a)
	// a[5] = 1
	a = append(a, 1)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/**
slice的 零值是 nil
一个nil 的slice的长度和容量都是0
*/

func NilSlice() {
	var z []int
	printSlice("z", z)

	if z == nil {
		fmt.Println("nil")
	}
}

/**
https://blog.golang.org/go-slices-usage-and-internals
append slice
*/

func Append() {
	var a []int
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}
