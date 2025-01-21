package main

import "fmt"

func function4() {
	fmt.Println("--- Trying to recover the function ---")
	if r := recover(); r != nil {

		fmt.Println("--- Program Recovered ---")

	}
}

func function3(s1, s2 int) bool {
	fmt.Println("Calc student grade")
	defer function4()

	grade := (s1 + s2) / 2

	if grade > 5 {
		return true
	} else if grade < 5 {
		return false
	}

	panic("Grade is equal to 5 !!!") // Panic clause kill the running program, but executes the defer functions before die

}

func main() {
	fmt.Println("-- Function Panic and Recover --")
	fmt.Println(function3(5, 5))
}
