package main

import "fmt"

func main() {

	// Basic Operators
	sum := 1 + 2
	sub := 2 - 1
	div := 4 / 2
	rest := 5 % 2
	mult := 3 * 5

	fmt.Println(sum, sub, div, rest, mult)

	// GO does no accept operations with varaibles of different types
	var num1 int8
	var num2 int16
	//var soma int16 = num1 + num2
	//fmt.Println(soma)

	fmt.Println(num1, num2)

	// End Basic Operators

	//Attribution
	var String1 string = "String 1 Attribution"
	String2 := "String 2 Attribution"

	fmt.Println(String1, String2)
	//End Attributions

	// Relational Operators
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	//End of relational operators

	// Logical Operators
	fmt.Println(true && false)
	fmt.Println(true || false)

	verdadeiro := true
	fmt.Println(!verdadeiro)
	// Fim Logical Operators

	// Unique Operators
	number1 := 10
	fmt.Println(number1)
	number1++ // Works only with + and - operations
	fmt.Println(number1)
	number1--
	fmt.Println(number1)
	number1 += 20
	fmt.Println(number1)
	number1 -= 2
	fmt.Println(number1)
	number1 -= 2
	fmt.Println(number1)

	number1 *= 2
	fmt.Println(number1)
	number1 /= 2
	fmt.Println(number1)
	number1 %= 3
	fmt.Println(number1)

}
