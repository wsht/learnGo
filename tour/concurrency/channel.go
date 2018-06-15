package concurrency

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int, step int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	sleepMill := time.Duration(1000*step) * time.Millisecond
	time.Sleep(sleepMill)
	fmt.Println(step)
	c <- sum
}

func Channel() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(a[:len(a)/2], c, 1)
	go sum(a[len(a)/2:], c, 2)

	t1 := <-c
	fmt.Println(t1)
	t2 := <-c
	fmt.Println(t2)

	// x, y := <-c, <-c
	fmt.Println(t1, t2, t1+t2)
}

/**
缓冲 channel
channel 可以是 _带缓冲的_。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：

ch := make(chan int, 100)
向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞。

修改例子使得缓冲区被填满，然后看看会发生什么。
*/

func addToChannel(timer int, c chan int) {
	for i := 0; i < timer; i++ {
		c <- i
		fmt.Printf("%d into chan at time %v\n", i, time.Now())
	}
	close(c)
}

func BufferedChannels() {
	c := make(chan int, 2)
	go addToChannel(100, c)
	for i := range c {
		fmt.Printf("%d out chan \n", i)
		time.Sleep(100 * time.Millisecond)
	}

}
