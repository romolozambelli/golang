package main

import "fmt"

func sumNumbers(numbers ...int) int {

	total := 0

	for _, numbers := range numbers {
		total += numbers
	}
	return total

}

// Variable functions accept only 1 different type parameter
func write(text string, numbers ...int) {
	for _, numbers := range numbers {
		fmt.Println("Numbers", numbers)
	}
}

func main() {
	total := sumNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println("Numbers", total)

	write("numero", 1, 2, 4)

}
