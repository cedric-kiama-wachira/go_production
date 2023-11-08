package main

import (
	"fmt"
)

func main() {
	var ckw string = "Cedric Kiama Wachira"
	if ckw == "Cedric Kiama Wachira" {
		fmt.Println(ckw)
	} else if ckw == "Cedric" {
		fmt.Println(ckw)

	} else if ckw == "Kiama Wachira" {
		fmt.Println("Only Two names given")
	} else {
		fmt.Println("None")
	}
}
