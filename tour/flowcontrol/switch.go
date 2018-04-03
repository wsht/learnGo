package flowcontrol

import (
	"fmt"
	"runtime"
	"time"
)

// 除非以fallthrough语句结束，否则分之会自动终止

func Switch() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("OS x")
		fallthrough
	default:
		fmt.Printf("%s.", os)
	}
}

func SwitchWithFallThrough() {
	switch i := 2; i {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		fallthrough
	case 3:
		fmt.Println(3)
		fallthrough

	default:
		fmt.Println(4)
	}
}

/**
switech 的条件是从上倒下执行的，当匹配成功时候停止
switch i{
case 0:
case f():
}

当 i==0时候 是不会调用`f`
*/

/**
没有条件的switch
没有条件的switch同switch true一样
这一构造是的可以用更清晰的形式来编写长的if-then-else
*/

func SwitchWithoutCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
