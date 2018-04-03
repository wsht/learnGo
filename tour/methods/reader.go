package methods

import (
	"fmt"
	"io"
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
