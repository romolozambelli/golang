package shapes

import (
	"math"
)

// Forma of a shape
type Forma interface {
	area() float64
}

// Properties of square
type Square struct {
	Lenght float64
	Weight float64
}

// Properties of a Circle
type Circle struct {
	raio float64
}

// Area of a Shape
func (r Square) Area() float64 {
	return r.Lenght * r.Weight

}

// Calculate the Area
func (r Circle) Area() float64 {
	return math.Pi * math.Pow(r.raio, 2)
}
