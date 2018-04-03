package flowcontrol

import (
	"fmt"
)

/**
defer 语句会延迟函数的执行，知道上层函数返回
延迟调用的参数会立即生成，但是在上层函数返回前都不会被调用
*/

func Defer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
