//Square struct with two fields: center of type point and length of type int.
//Add two methods Move and Area for calculation

package main

import "fmt"

//Point is a 2D point
type Point struct {
	X int //x coordinate
	Y int //y coordinate
}

//Move moves the point
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

//Square struct
type Square struct {
	Length int
}

//creates a new square
func NewSquare(length int) (*Square, error) {
	return &Square{
		Length: length,
	}, nil
}

//returns the area of the square
func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {

	fmt.Println("=========Start==========")
	square, err := NewSquare(5)

	if err == nil {
		fmt.Println(square.Area())
	}

	fmt.Println("=========End==========")

}
