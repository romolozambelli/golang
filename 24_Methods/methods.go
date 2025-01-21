package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func write() {
	fmt.Println("Writing on the shell")
}

func (u user) save() {
	fmt.Println("Save the number")
	fmt.Printf("Persisiting the data of %s \n", u.name)
}

func (u user) allowDrink() bool {
	// if u.age > 17 {
	// 	return true
	// } else {
	// 	return false
	// }
	return u.age > 17
}

func (u *user) birthday() {
	u.age++
}

func main() {
	user1 := user{"Jose", 23}
	fmt.Println(user1)
	user1.save()

	user2 := user{"David", 24}
	user2.save()
	fmt.Println(user2.allowDrink())

	user3 := user{"Frank", 17}
	fmt.Println("---Birthday---")
	fmt.Println(user3.age)
	user3.birthday()
	fmt.Println(user3.age)

}
