package main

import (
	"fmt"
)

func main() {
	fmt.Println("Internal Arryays")

	slice1 := make([]float32, 10, 11)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, 13)
	fmt.Println(slice1)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Appending more items than supported by a slice. Go always allocate more memory than is necessary
	slice1 = append(slice1, 13)
	fmt.Println(slice1)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := make([]float32, 5)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice2 = append(slice2, 12)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}
