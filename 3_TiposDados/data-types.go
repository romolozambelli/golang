package main

import (
	"errors"
	"fmt"
)

func main() {

	// BEGIN INTEGER

	//int8, int16, int32, int64
	//int uses the computer arch to define the int size
	var number_1 int8 = 100

	fmt.Println(number_1)

	//Alias
	// RUNE = Int32. In general used to reference variables that reflects characteres from ASC table

	//BYTE = type int8

	//uint are non negative numbers
	//Follows the same size as the int numbers

	var number_2 uint8 = 200
	fmt.Println(number_2)

	//float32 and float64 accepts broken numbers

	var number_3 float32 = 20.2
	fmt.Println(number_3)

	var number_4 float64 = 20.45
	fmt.Println(number_4)

	// Initial value is equal to "0"
	var number_5 int
	fmt.Println(number_5)

	// END NUMBERS

	// BEGIN STRINGS

	var stringVariable_1 string = "String_1"
	fmt.Println(stringVariable_1)

	stringVariable_2 := "String_2"
	fmt.Println(stringVariable_2)

	character_1 := 'R'
	fmt.Println(character_1)

	character_2 := 'O'
	fmt.Println(character_2)

	// Initial value of strings is empty
	var stringVariable_3 string
	fmt.Println(stringVariable_3)

	// END STRINGS

	// BEGIN BOOLEAN
	var boolean_1 bool = true
	fmt.Println(boolean_1)

	var boolean_2 bool = false
	fmt.Println(boolean_2)

	var boolean_3 bool
	fmt.Println(boolean_3)
	//END BOOLEAN

	//BEGIN ERROR
	var erro_1 error
	fmt.Println(erro_1)

	var erro_2 error = errors.New("Internal failure !!!")
	fmt.Println(erro_2)

	// END ERRORS
}
