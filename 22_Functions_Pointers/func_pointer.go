package main

import "fmt"

func inverterNumber(n1 int) int {
	return n1 * -1
}

func inverterNumberPointer(n1 *int) {
	*n1 = *n1 * -1
}

func main() {
	number := 20

	converterNumber := inverterNumber(number)

	fmt.Println(converterNumber) // Manipulate the value not the value

	inverterNumberPointer(&number) // Manipulate the memory position

	fmt.Println(number)

	newNumber := 40

	inverterNumberPointer(&newNumber)
	fmt.Println(newNumber)

}
