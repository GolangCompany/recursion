package main

import (
	"fmt"
)

func tail(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	tail(n - 1)
}

func main() {
	tail(3)
}
