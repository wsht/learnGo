package moretypes

import (
	"fmt"
)

type Vertex struct {
	x int //小写的时候 外部不能引用以及赋值
	Y int
}

func Struct() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(v.x)
	fmt.Println(v.Y)
}

/**
结构体指针
*/
func StructPointer() {
	v := Vertex{1, 2}
	p := &v
	p.x = 1e9
	fmt.Println(v)
}

/**
结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。
*/
func StructLiterals() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{x: 1}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)

	fmt.Println(v1, p, v2, v3)
	fmt.Println(*p)
}
