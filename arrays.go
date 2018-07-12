//An array is a numbered sequenced of elements of a single type with a fixed length.
//Declaration : var name [size]type

package main

import "fmt"

func main() {
	var x [5]int
	x[4]=100
	fmt.Println(x)

	var  str [10]string
	str[1]="Sandeep"
	fmt.Println(str) //if you look at the output it will have blanks for empty slots

	var intArr [10]int
	for i:=0;i<10;i++ {
		intArr[i]=i
	}
	fmt.Println(intArr)
	fmt.Println("Length of intArr is : ", len(intArr))

	var floatArr [10]float64
	total:=0.0
	for i:=0;i<10;i++ {
		floatArr[i]=float64(i) * 0.2 //to convert between types, we use the type name as a function
		total+=floatArr[i]
	}
	fmt.Println(floatArr)
	fmt.Println("Total of all the elements : ", total)

	//short hand syntax for creating arrays
	shortArr := [10]int {10,20,30,40,50,60,70,80,90,100}//, "S"} - try to add another type and see compilation error
							  //,110} //similarly try to add one more element in the declaration and see compilation erro
	fmt.Println("shortArr", shortArr)

	total = 0.0
	//Go's version of for each
	for _, value:=range shortArr {
		total+=float64(value)
	}

	fmt.Println("Average : ", total/float64(len(shortArr)))

	//dynamic arrays

	//size:=15 //uncomment this to see what happens when you try to create a array with variable size
	const size=15
	var dynArr [size]int
	for i:=0; i<size;i++ {
		dynArr[i]=i
	}

	fmt.Println("dynamic array", dynArr)
}
