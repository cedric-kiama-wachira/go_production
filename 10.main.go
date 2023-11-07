package main

import (
	"fmt"
	"reflect"
)

func main() {

	var type_int int = 9999999
	var type_f32 float32 = 99.99
	var type_str string = "Cedric"
	var type_t_or_f bool = true

	// This section will use the %T format specifier
	fmt.Println("Using Type T format Identifier Results Start Here")
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("The variable type_integer = %v is of type %T\n", type_int, type_int)
	fmt.Printf("The variable type_float32 = %v is of type %T\n", type_f32, type_f32)
	fmt.Printf("The variable type_string = '%v' is of type %T\n", type_str, type_str)
	fmt.Printf("The variable type_t_or_f = '%v' is of type %T\n", type_t_or_f, type_t_or_f)
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Using Using the Reflect Package and Type Of Function starts Here")
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("The Variable name is type_int, and the Data Type is: %v\n", reflect.TypeOf(type_int))
	fmt.Printf("The Variable name is type_f32, and the Data Type is:%v\n", reflect.TypeOf(type_f32))
	fmt.Printf("The Variable name is type_str, and the Data Type is: %v\n", reflect.TypeOf(type_str))
	fmt.Printf("The Variable name is type_t_or_f, and the Data Type is: %v\n", reflect.TypeOf(type_t_or_f))
	fmt.Println("-------------------------------------------------------")
}
