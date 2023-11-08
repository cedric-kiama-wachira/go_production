package main

import (
	"fmt"
)

func main() {
	var (
		a, b int = 200, 20
	)
	a /= b
	fmt.Println(a)
}
