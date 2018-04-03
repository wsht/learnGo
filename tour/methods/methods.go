package methods

import (
	"fmt"
	"math"
)

/**
Go 没有类，然而仍然可以在结构类型上定义方法
方法接受者出现在 func关键字和方法名之间的参数中
*/

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}

/**
你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。

但是，不能对来自其他包的类型或基础类型定义方法。
*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func MethodsContinued() {
	f := MyFloat(-math.Sqrt2)

	fmt.Println(f.Abs())
}

/**

接收者为指针的方法
方法可以与命名类型或命名类型的指针关联。

刚刚看到的两个 Abs 方法。一个是在 *Vertex 指针类型上，而另一个在 MyFloat 值类型上。 有两个原因需要使用指针接收者。首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。其次，方法可以修改接收者指向的值。

尝试修改 Abs 的定义，同时 Scale 方法使用 Vertex 代替 *Vertex 作为接收者。

当 v 是 Vertex 的时候 Scale 方法没有任何作用。`Scale` 修改 `v`。当 v 是一个值（非指针），方法看到的是 Vertex 的副本，并且无法修改原始值。

Abs 的工作方式是一样的。只不过，仅仅读取 `v`。所以读取的是原始值（通过指针）还是那个值的副本并没有关系
*/

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodsWithPointerReceivers() {
	v1 := &Vertex{3, 4}
	v2 := Vertex{3, 4} /*notice v2是否是指针类型 和Scale方法中的指针没有必要关系？*/ /*TODO*/

	v1.Scale(2)
	fmt.Println(v1, v1.Abs())
	// v1.Scale2(2)
	// fmt.Println(v1, v1.Abs2())

	v2.Scale(2)
	fmt.Println(v2, v2.Abs())
	// v2.Scale2(2)
	// fmt.Println(v2, v2.Abs2())
}

/**

接口
接口类型是由一组方法定义的集合。

接口类型的值可以存放实现这些方法的任何值。
*/

type Abser interface {
	Abs() float64
}

type Abser2 interface {
	Abs2() float64
}

func Interface() {
	var a Abser
	var a2 Abser2

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v

	a2 = v

	fmt.Println(a.Abs())
	fmt.Println(a2.Abs2())
}
