package main

import "fmt"

func closure() func() {
	text := "-- Function Closure --"

	function1 := func() {
		fmt.Println(text)
	}
	return function1
}

func main() {
	text := "Inside of Function Main"
	fmt.Println(text)

	newFunction := closure()
	newFunction()

}
