//Slice is a segment of an array. It has a length but it is changeable. 
//Declaration : var s []int. The only difference between slice and array declaration is the missing length
//make function is used to create a slice

package main

import "fmt"

func main() {
	var x []int
	x=make([]int, 5)
	fmt.Println(x)

	x=make([]int, 5, 10) //slice with length 5 and underlying array of capacity 10
	fmt.Println(x)

	y:=[10]int {0,1,2,3,4,5,6,7,8,9}
	x=y[0:5] //slicing an array from its indices
	fmt.Println(x)
	
	x=y[5:] //slicing given start param only. 
	fmt.Println(x)

	x=y[:5] //slicing given end param only. Notice here 5th index is excluded
	fmt.Println(x)

	x=y[:]
	fmt.Println(x)

	x=y[0:]
	fmt.Println(x)

	x=y[0:len(y)]
	fmt.Println(x)

	//append(slice, ...elements,...) is the syntax
	x=append(x, 10,11,12) //append can add elements to a slice
	fmt.Println(x,y)

	z:= make([]int ,5)
	//copy(dest, src) is the syntax
	copy(x,z) //copy will copy the elements from src slice to dest slice without changing the size of any slice. See print output. z has only capacity of 5
	fmt.Println(x, z)
}
