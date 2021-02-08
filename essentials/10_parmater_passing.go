package main

import "fmt"

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	fmt.Println("============Start=============")

	values := []int{1, 2, 3, 4}
	doubleAt(values, 2) //this is passed by reference, so any changes inside the function stay after the function is over
	fmt.Println(values)

	val := 10
	double(val) //this is passed by value, see there will be no change
	fmt.Println(val)
	doublePtr(&val)
	fmt.Println(val)

	fmt.Println("============End=============")
}
