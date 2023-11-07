package main

// Converting Variables from One Data Type to another - This process is known as Type Casting
import (
	"fmt"
	"strconv"
)

func main() {
	var first_int_no int = 90
	var first_float_no float64 = float64(first_int_no)
	var second_float_no float64 = 90.76
	var second_int_no int = int(second_float_no)
	var string_conversion string = strconv.Itoa(first_int_no)

	fmt.Println("Results of the Integer Value 90 addding six decimal points")
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("%.6f\n", first_float_no)
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Results of the Float Value 90.76 but truncating the decimal points")
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("%v\n", second_int_no)
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("Introducing the use of the STRCONV package to one, convert integers to string ")
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Printf("%q\n", string_conversion)
	fmt.Println("------------------------------------------------------------------------------")
}
