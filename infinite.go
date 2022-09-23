package main

import (
	"fmt"
)

func infinite() {
	fmt.Println("Welcome to Golang.Company")
	infinite()
}

func main() {
	infinite()
}
