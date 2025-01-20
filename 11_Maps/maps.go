package main

import "fmt"

func main() {
	fmt.Println("Maps ")

	// Keys and Values should have the same value on Maps
	user := map[string]string{
		"name":    "ronald",
		"surname": "mcdonalds",
	}
	fmt.Println(user)

	user2 := map[string]map[string]string{
		"name": {
			"first":  "Arthur",
			"second": "King",
		},
		"addrees": {
			"roda":   "street 1",
			"number": "number 1",
		},
	}

	fmt.Println("Before Deleting Keys")
	fmt.Println(user2)

	fmt.Println("Deleting Keys")
	delete(user2, "name")
	fmt.Println(user2)
}
