package main

import (
	"fmt"
)

func factorial(x int) int {
	if x == 1 || x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(factorial(x))
}
