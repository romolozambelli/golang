package main

import "fmt"

func returnNumberFibonacci(position uint) uint {
	if position <= 1 {
		return position

	}

	return returnNumberFibonacci(position-2) + returnNumberFibonacci(position-1)
}

func main() {
	fmt.Println("-- Recursive Functions --")

	position := uint(3)

	for i := uint(0); i < position; i++ {
		fmt.Println(returnNumberFibonacci(i))
	}
}
