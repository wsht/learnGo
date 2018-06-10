package main

import (
	"fmt"
	"testing"
	"time"
)

type Pool struct {
	Queue          chan func() error
	RuntineNumber  int
	Total          int
	Result         chan error
	finishCallback func()
}

func (self *Pool) Init(runtineNumber int, total int) {
	self.RuntineNumber = runtineNumber
	self.Total = total
	self.Queue = make(chan func() error, total)
	self.Result = make(chan error, total)
}

func (self *Pool) Start() {
	//开启number个goruntine
	for i := 0; i < self.RuntineNumber; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}
				err := task()
				self.Result <- err
			}
		}()
	}
	//获取每个任务的处理结果
	for j := 0; j < self.Total; j++ {
		res, ok := <-self.Result
		if !ok {
			break
		}
		if res != nil {
			fmt.Println(res)
		}
	}

	//结束回调函数
	if self.finishCallback != nil {
		self.finishCallback()
	}
}

func (self *Pool) Stop() {
	close(self.Queue)
	close(self.Result)
}

func (self *Pool) AddTask(task func() error) {
	self.Queue <- task
}

func (self *Pool) SetFinishCallback(f func()) {
	self.finishCallback = f
}

func Download(url string) error {
	time.Sleep(1 * time.Second)
	fmt.Println("Downlaod " + url)
	return nil
}

func DownloadFinish() {
	fmt.Println("Download finsh")
}

func TestPool(t *testing.T) {
	var p Pool
	url := []string{"11111", "22222", "33333", "444444", "55555", "66666", "77777", "88888", "999999"}
	p.Init(3, len(url))

	for i := range url {
		u := url[i]
		p.AddTask(func() error {
			return Download(u)
		})
	}

	p.SetFinishCallback(DownloadFinish)
	p.Start()
	p.Stop()
}
