package main

import (
	"fmt"
)

func main() {

	i := 0

	for i < 10 {
		i++
		//time.Sleep(time.Second)
		fmt.Println(i)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Increasing J = ", j)
		//time.Sleep(time.Second)
	}

	// fmt.Println(j) - The j variable is only part of the if scope

	names := [3]string{"john", "bob", "adam"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for index, letter := range "word" {
		fmt.Println(index, letter) // The letter are printing the ASC tables numbers

		fmt.Println(index, string(letter)) // The letter are printing the ASC tables numbers
	}

	usermap := map[string]string{
		"name":    "john",
		"surname": "Silva",
	}

	for key, value := range usermap {
		fmt.Println(key, value)
	}

	// for{
	// infinite loop
	// }

	// GOLang does not accept for on struct
	//type userStruct struct{"Ze","Colmeia"}

	// for key, value := range usermap {
	// 	fmt.Println(key, value)
	// }

}
