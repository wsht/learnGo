package goroutineController

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const N = 10

/**
利用channel 特性，来实现监控所有的goroutine是否完成
*/
func TestWithoutWaitGroup(test *testing.T) {
	waitChan := make(chan int, 1)
	for i := 0; i < N; i++ {
		go func(n int) {
			fmt.Println(n)
			waitChan <- 1
		}(i)
	}

	cnt := 0
	for range waitChan {
		cnt++
		if cnt == N {
			break
		}
	}

	close(waitChan)
	fmt.Println("finished")
}

/**
可以通过 sync.WaitGroup 来等待一个goroutines集合的结束
*/
func TestWaitGroup(test *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	wg.Wait()
	fmt.Println("finished")
}

/**
如何限制goroutine的创建数量(信号量实现)
*/

type Semaphore struct {
	Threads chan int
	Wg      sync.WaitGroup
}

func NewSemaphore(n int) *Semaphore {
	inst := new(Semaphore)
	inst.Threads = make(chan int, n)
	return inst
}

/**
资源申请
*/
func (sem *Semaphore) P() {
	sem.Threads <- 1
	sem.Wg.Add(1)
}

/**
资源释放
*/
func (sem *Semaphore) V() {
	sem.Wg.Done()
	<-sem.Threads
}

//wait for all unit resources to be released
func (sem *Semaphore) wait() {
	sem.Wg.Wait()
}

func TestSemaphore(t *testing.T) {
	s := NewSemaphore(5)

	for index := 0; index < 20; index++ {
		s.P()
		go func() {
			r := rand.Intn(3)
			fmt.Println("sub called, sleeping", r)
			// t.Log("sub called, sleeping", r)
			time.Sleep(time.Duration(r) * time.Second)
			s.V()
		}()
	}
}

/**
如何让goroutine主动退出
*/

func TestGoroutineQuit(test *testing.T) {
	ok, quit := make(chan int, 1), make(chan int, 1)
	go func() {
		i := 0
		for {
			select {
			case <-quit:
				ok <- 1
				return
			default:
				fmt.Println(i)
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	quit <- 1
	<-ok
}
