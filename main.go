package main

import (
	"fmt"
)

func main() {
	var first_name_of_the_multi_line string
	var second_name_of_the_multi_line string
	var third_name_of_the_multi_line string
	var are_this_your_full_names bool = true

	fmt.Println("Enter your first name: ")
	fmt.Scanf("%s", &first_name_of_the_multi_line)
	fmt.Println("Enter your second name: ")
	fmt.Scanf("%s", &second_name_of_the_multi_line)
	fmt.Println("Enter your third name: ")
	fmt.Scanf("%s", &third_name_of_the_multi_line)
	fmt.Println("Confirm with either true or false to see results. ")
	fmt.Scanf("%t", &are_this_your_full_names)
	fmt.Println(first_name_of_the_multi_line, second_name_of_the_multi_line, third_name_of_the_multi_line, are_this_your_full_names)
}
