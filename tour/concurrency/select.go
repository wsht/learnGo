package concurrency

import (
	"fmt"
	"time"
)

/**
select语句使得一个goroutine 在多个通讯操作上等待
select 会堵塞，知道条件分支中的某个可以继续执行，这时就会执行那个条件分支
当多个都准备好的时候，会随机选择一个(??随机，这个有点坑)
*/

func fibonacci_test(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func SelectGo() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci_test(c, quit)
}

func DefaultSelect() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("      .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
