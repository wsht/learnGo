package moretypes

import "fmt"

/**
for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。
可以通过赋值给 _ 来忽略序号和值。

如果只需要索引值，去掉“, value”的部分即可。
*/

func Range() {
	var array = [10]int{1, 2, 3, 4, 5}

	var slice = []int{1, 2, 3, 4, 5}

	for index, value := range array {
		fmt.Println(index, value)
	}

	for index, value := range slice {
		fmt.Println(index, value)
	}

}
