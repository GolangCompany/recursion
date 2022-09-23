package main

import (
	"fmt"
)

func main() {
	var anon func(int)
	anon = func(x int) {
		if x == 0 {
			return
		}

		fmt.Println(x)
		anon(x - 1)

	}

	anon(4)
}
