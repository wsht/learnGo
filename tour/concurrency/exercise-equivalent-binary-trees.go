package concurrency

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int, notRoot bool) {
	ch <- t.Value
	left := t.Left
	if left != nil {
		// fmt.Println(t.Value, "left")
		// fmt.Println(left.Value)
		Walk(left, ch, true)
	}

	right := t.Right
	if right != nil {
		// fmt.Println(t.Value, "right")
		// fmt.Println(right.Value)s
		Walk(right, ch, true)
	}

	if !notRoot {
		close(ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	data1 := make([]int, 10)
	data2 := make([]int, 10)

	go Walk(t1, ch1, false)
	go Walk(t2, ch2, false)

	for value := range ch1 {
		fmt.Println("ch1\n")
		data1 = append(data1, value)
	}

	fmt.Println("end ch1\n")
	for value := range ch2 {
		data2 = append(data2, value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(data1)))
	sort.Sort(sort.Reverse(sort.IntSlice(data2)))

	for index, value := range data1 {
		if value != data2[index] {
			return false
		}
	}

	return true
}

func ExerciseEquivalentBinaryTree() {
	// t := tree.New(1)
	// fmt.Println(t)
	// c := make(chan int)

	// Walk(t, c)

	// for v := range c {
	// 	// fmt.Println(v)
	// }
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
