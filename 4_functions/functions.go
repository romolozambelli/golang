package main

import "fmt"

func sum(a int8, b int8) int8 {
	return a + b
}

// go supports returning two varaibles
func calcMaths(n1, n2 int8) (int8, int8) {

	soma := n1 + n2
	minus := n2 - n1

	return soma, minus
}

func main() {

	total := sum(10, 20)
	fmt.Println(total)

	// go supports variables with func types
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Func receives strings")
	fmt.Println(result)

	fmt.Println(calcMaths(10, 5))

	resultSum, resultSub := calcMaths(30, 10)
	fmt.Println(resultSum, resultSub)
}
