package basic

import "fmt"

const Pi = 3.14

const Pi2 = 10 + 2

func Constants() {
	const World = "world"
	fmt.Println("happy", Pi, "days")
	fmt.Println("hello", World)
	const Truth = true

	fmt.Println("go rules", Truth)

	result := 2 * Pi2
	fmt.Println(result)
}
