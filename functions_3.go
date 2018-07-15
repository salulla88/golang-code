//By default, the parameters in functions are passed as values. i.e. you need to always return a value after operation from the function and then use it.
//But what if you do not want to do return values, instead you want to pass parameters as reference. For this, Go supports pointers
//* and & are the same syntax from the C times for pointers in Go as well.
//In addition, if you would like to get a new pointer, you can, by using "new" keyword

package main

import "fmt"

func main() {
	total := 100
	
	changeTotal(total)
	fmt.Println("total after calling changeTotal : ", total)
	
	changeTotalByReference(&total)
	fmt.Println("total after calling changeTotalByReference", total)

	x := new(int)
	assignValues(x)
	fmt.Println("Values of x post assignment : ", *x)
}

//By default, parameters are passed as values
func changeTotal(total int) {
	total = total * 5
}

//If we need to change the value of the parameter being passed, we need to pass it as reference i.e. pointer to that variable
func changeTotalByReference(pointerToTotal *int) {
	*pointerToTotal = *pointerToTotal * 5
}

func assignValues(ptr *int) {
	*ptr = 500
}
