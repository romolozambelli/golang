package main

import "fmt"

func main() {

	type address struct {
		street string
		number int8
	}

	type user struct {
		username string
		name     string
		password string
		age      int8
		address  address
	}

	var u user
	u.username = "Anom"
	u.name = "John"
	u.password = "12345"
	u.age = 32

	addressSample := address{"Rua asfdsdf", 10}

	fmt.Println(u)

	user2 := user{"username", "andrew", "123124", 43, addressSample}
	fmt.Println(user2)

	user3 := user{name: "testName"}

	fmt.Println(user3)

}
