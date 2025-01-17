package main

import "fmt"

// Go does not support inheritance by defaulkt since it is not object oriented language

type person struct {
	name      string
	surnamece string
	age       int8
}

type student struct {
	person
	school string
	course string
}

func main() {
	fmt.Println("Inheritance")

	p1 := person{"peter", "crowley", 24}

	fmt.Println(p1)

	e1 := student{p1, "private school", "computer science"}

	fmt.Println(e1)
	fmt.Println(e1.name)

}
