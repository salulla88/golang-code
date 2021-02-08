package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) { //its important to give pointer to the struct
	p.X += dx
	p.Y += dy
}
func main() {
	fmt.Println("=========Start=========")

	p := &Point{1, 2}
	p.Move(2, 3)
	fmt.Printf("%+v\n", p)
	fmt.Println("=========End=========")
}
