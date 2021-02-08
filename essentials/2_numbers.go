//Calculate mean of two numbers
package main

import (
	"fmt"
)

func main() {
	fmt.Println("==============Start=============")

	var x int
	var y int

	fmt.Println("Before assignment ", x, y)

	x = 1
	y = 2

	fmt.Println("After assignment ", x, y)

	//	var mean int
	fmt.Println("Mean : ", (x+y)/2)

	var meanInt int
	var meanFloat float64
	meanInt = (x + y) / 2
	meanFloat = float64(x+y) / 2
	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("meanInt=%v, type of %T\n", meanInt, meanInt)
	fmt.Printf("meanFloat=%v, type of %T\n", meanFloat, meanFloat)

}
