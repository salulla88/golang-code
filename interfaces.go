//In Go, interfaces are collections of method signatures
//Like Java or other programming languages, we do not have a extends or implements keyword
//We just implement the same function for different types as defined in interface and Go understands it.
//Structs should implement all the methods of an interface else none.

package main

import "fmt"

//Shape is an interface that has an area method.
type Shape interface {
	area() float64
//	perimeter() float64 //uncomment this line and see what error you get if you  do not implement that for your struct types
}

type Square struct {
	length float64
}

type Rectangle struct {
	length, width float64
}

//Implement area for square shapes
func (s *Square) area() float64 {
	return s.length * s.length
}

//Implement area for rectangle shapes
func (r *Rectangle) area() float64 {
	return r.length * r.width
}

func calculateArea(s Shape) {
	fmt.Println(s)
	fmt.Println("Area : ", s.area())
}

func main() {
	s := new(Square)
	s.length = 10.9
	calculateArea(s)

	r := new(Rectangle)
	r.length = 10.9
	r.width = 11.10
	calculateArea(r)
}

