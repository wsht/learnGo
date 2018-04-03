package flowcontrol

import (
	"fmt"
	"io"
	"os"
)

/**
defer 栈
延迟的函数调用被压入一个栈中。当函数返回时，会按照后进先出的顺序调用被延迟的函数调用。
*/

func DeferMulti() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

/**
https://blog.golang.org/defer-panic-and-recover
*/

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
