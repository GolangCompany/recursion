package main

import (
	"fmt"
)

func head(n int) {
	if n == 0 {
		return
	}
	head(n - 1)
	fmt.Println(n)
}
func main() {
	head(3)
}
