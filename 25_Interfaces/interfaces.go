package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type square struct {
	lenght float64
	weight float64
}

type circle struct {
	raio float64
}

func (r square) area() float64 {
	return r.lenght * r.weight

}

func (r circle) area() float64 {
	return math.Pi * math.Pow(r.raio, 2)
}

func printArea(f forma) {
	fmt.Println("The area of the shape ", f.area())
}

func main() {
	r := square{10, 15}
	printArea(r)

	c := circle{10}
	printArea(c)

}
