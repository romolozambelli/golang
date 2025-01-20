package main

import "fmt"

func mathCalc(n1, n2 int) (sum int, sub int) {

	sum = n1 + n2
	sub = n1 - n2
	return

}

func main() {

	fmt.Println(mathCalc(10, 20))

}
