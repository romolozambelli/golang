package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello World")
	}()

	func(text string) {
		fmt.Println(text)
	}("Welcome to the world")

	retorno := func(text string) string {
		return fmt.Sprintf("Received String -> %s Case", text)
	}("Welcome to the new world")

	fmt.Println(retorno)
}
