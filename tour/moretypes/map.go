package moretypes

import (
	"fmt"
)

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func Map() {
	m = make(map[string]Vertex2)
	m["bell labs"] = Vertex2{
		40.123,
		-74.123,
	}

	fmt.Println(m["bell labs"])

	t := make(map[string]int)

	t["hello"] = 1
	t["world"] = 2

	fmt.Println(t)

	a := make(map[int][]int)
	a[1] = []int{1, 2, 3}
	a[2] = []int{4, 5, 6}
	a[3] = []int{7, 8, 9}

	b := a[3]

	b[1] = 10

	fmt.Println(a)

}

func MutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
