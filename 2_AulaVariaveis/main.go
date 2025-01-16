package main

import "fmt"

func main() {
	var stringvar1 string = "StringVar_1"
	stringvar2 := "StringVar_2"

	fmt.Println(stringvar1)
	fmt.Println(stringvar2)

	var (
		stringvar3 string = "StringVar_3"
		stringvar4 string = "StringVar_4"
	)

	fmt.Println(stringvar3, stringvar4)

	stringvar5, stringvar6 := "StringVar_5", "StringVar_6"

	fmt.Println(stringvar5, stringvar6)

	const constante1 string = "constant_1"

	fmt.Println(constante1)

	stringvar5, stringvar6 = stringvar6, stringvar5

	fmt.Println(stringvar5, stringvar6)

}
