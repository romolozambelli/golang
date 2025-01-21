package main

import "fmt"

func function1() {
	fmt.Println("Running Function 1")
}

func function2() {
	fmt.Println("Running Function 2")
}

func function3(s1, s2 int) bool {
	defer fmt.Println("Student Grade Ready") // Executed before the return ...
	fmt.Println("Calc student grade")

	grade := (s1 + s2) / 2

	if grade > 5 {
		//fmt.Println("Student Grade Ready")
		return true
	}
	//fmt.Println("Student Grade Ready")
	return false
}

func main() {
	//defer function1()
	//function2()
	fmt.Println(function3(1, 4))
}
