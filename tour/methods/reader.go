package methods

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/reader"
)

func ReaderTest() {
	r := strings.NewReader("hello, Reader")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct {
}

var str = strings.Repeat("A", 1000)
var strlen = len(str)
var already = 0

func (r MyReader) Read(p []byte) (int, error) {
	var lens int
	remainingLens := strlen - already
	if remainingLens > len(p) {
		lens = len(p)
	} else {
		lens = remainingLens
	}

	// pointer := &p

	for i := 0; i < lens; i++ {
		p[i] = str[i]
	}

	return remainingLens, nil
}

func ExerciseReader() {
	reader.Validate(MyReader{})
}

/**
练习：rot13Reader
一个常见模式是 io.Reader 包裹另一个 `io.Reader`，然后通过某种形式修改数据流。

例如，gzip.NewReader 函数接受 `io.Reader`（压缩的数据流）并且返回同样实现了 io.Reader 的 `*gzip.Reader`（解压缩后的数据流）。

编写一个实现了 io.Reader 的 `rot13Reader`， 并从一个 io.Reader 读取， 利用 rot13 代换密码对数据流进行修改。

已经帮你构造了 rot13Reader 类型。 通过实现 Read 方法使其匹配 `io.Reader`。
*/

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (int, error) {
	//read from r.r
	//exchage to rot13
	//results
	n, err := r.r.Read(p)
	for index, value := range p {
		p[index] = value + 13
	}

	return n, err
}

func ExerciseRotReader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
