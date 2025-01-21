package main

import "fmt"

var n int

func init() {
	fmt.Println("Function executed before the main")
	n = 10
}

func main() {
	fmt.Println("Program Beggining")
	fmt.Println(n)
}
