package main

import (
	"fmt"
)

var n int

func odd() {
	if n <= 10 {
		fmt.Println(n + 1)
		n++
		even()
	}
	return
}
func even() {
	if n <= 10 {
		fmt.Println(n - 1)
		n++
		odd()
	}
	return
}
func main() {
	n = 1
	odd()
}
