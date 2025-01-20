package main

import "fmt"

func main() {

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++

	fmt.Println(var1, var2)

	// Pointers are a memory reference

	var var3 int
	var pointer *int

	var3 = 100
	pointer = &var3 // Reference to a memory position

	fmt.Println(var3, pointer)
	fmt.Println(var3, *pointer) // Reasding position size value

	var3++
	fmt.Println(var3, pointer)
	fmt.Println(var3, *pointer)

}
