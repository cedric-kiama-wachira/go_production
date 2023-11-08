package main

import (
	"fmt"
)

func main() {
	var s int = 100
	switch s {
	case 10:
		fmt.Println("s is 10")
	case 100, 200:
		fmt.Println("s is either 100 or 200")
	default:
		fmt.Println("s i neither 10, 100 or 200")

	}
}
