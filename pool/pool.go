package main

type TpThreadPool struct {
	MinThNum int //线程池最小线程数
	curThNum int //线程池当前线程数量
	MaxThNum int //线程池最大线程数量

}

//how dynamic set goruntine num
