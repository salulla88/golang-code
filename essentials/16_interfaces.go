//interfaces
package main

import (
	"fmt"
	"math"
)

//Square
type Square struct {
	Length float64
}

//Area returns the area of the square
func (s *Square) Area() float64 {
	return s.Length * s.Length
}

//Circle
type Circle struct {
	Radius float64
}

//Area returns the area of the circle
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//sumAreas returns the sum of all areas in the slice
func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

type Shape interface {
	Area() float64
}

func main() {
	fmt.Println("===========Start===========")

	s := &Square{20}
	fmt.Println(s.Area())

	c := &Circle{10}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Println(sa)

	fmt.Println("===========End===========")
}
