package main

import (
	"fmt"
)

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Go testing")
	result := Calculate(2)
	fmt.Println(result)
}

func Add(x, y int) int {
	return x + y
}

// func main() {
// 	fmt.Println("Hello")
// }
