package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	var array1 [6]int8

	fmt.Println(array1)

	array1[2] = 2

	fmt.Println(array1)

	array2 := [3]string{"String1", "String2", "String3"}
	fmt.Println(array2)

	array3 := [...]string{"String1", "String2", "String3", "String4"}
	fmt.Println(array3)

	// Slice does not have a fixed size
	slice1 := []int{1, 2, 4}
	fmt.Println(slice1)
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	slice1 = append(slice1, 5)
	fmt.Println(slice1)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	fmt.Println(array3)
	array3[1] = "new value"
	fmt.Println(array3)

}
