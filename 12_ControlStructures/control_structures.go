package main

import "fmt"

func main() {

	fmt.Println("Control Structures")

	number1 := 20

	// GO does not accept if withouth the clauses
	if number1 > 15 {
		fmt.Println("higher than 15")

	} else {
		fmt.Println("Lower than 15")
	}

	if number2 := number1; number2 > 0 {
		fmt.Println("higher than 0")
		fmt.Println(number2)
	}

	// if init limits the scope of the variable to the conditions
}
