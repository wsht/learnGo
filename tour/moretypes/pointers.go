package moretypes

import "fmt"

/**
 go 具有指针，指针保存了变量的内存地址
 类型 *T 师只想类型T的值的指针。其零值是 `nil`

var p *int

& 符号会生成一个只想其作用对象的指针
i:=42
p = &i

* 符号表示指针指向的底层的值
fmt.println(*p) //通过指针p 读取i的值
*p = 21 //通过指针p 设置i

这也就是通常所说的间接引用

*/
func Pointers() {
	i, j := 42, 2701

	p := &i         //point to i
	fmt.Println(p)  //echo i pointer value
	fmt.Println(*p) //read  i through the pointer
	*p = 21
	fmt.Println(i) //see the new value of i

	p = &j         //point to j
	*p = *p / 37   //divide j through the pointer
	fmt.Println(j) //see the new value of j
}
