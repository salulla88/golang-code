//Closure is a concept long known to javascript developers. Go also supports closure.
//It is possible to create functions within functions.
//And the internal functions have access to the variables of the caller function.

package main

import "fmt"

func main() {
	z := 5
	result := func(x, y int) int {
		return (x + y) * z
	}

	fmt.Println(result(1,2))

	fmt.Println("Factorial of 0 : ", factorial(0))
	fmt.Println("Factorial of 9 : ", factorial(9))
}

//best example for a recursive function is factorial. Below is one in Go	
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}

//Closure and recursion form the basis of a paradigm called functional programming
