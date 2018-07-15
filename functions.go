//Unlike other languages, in Go, functions are sections of code that map zero or more input parameters to zero or more output parameters
//Declaration : func <funcName> (paramname paramtype, paramname, paramtype,...) <returntype>
//Functions form call stacks i.e the first caller is the last one to return.
//In Go, return types can have names e.g. func m1() (r int). u can set the value of r and just return.
//In Go, functions can return multiple values. e.g. func m1() (int, boolean)
//Go also provides option of passing variable parameters to functions. e.g. func m1(args ...int) int
package main

import "fmt"

func main() {
	fmt.Println("Hello to functions in Go!!!")

	marks:=[]int{58,98,62,74,91}
	fmt.Println("Average score : ", calculateAverage(marks))
	fmt.Println("Average score with named return type : ", namedReturnType(marks))
	avg, isFail := multipleReturnValues(marks)
	fmt.Println("Average with multiple return types : ", avg, isFail)
	fmt.Println("Average with variable parameters : ", variableLengthParameters(10,20,30,40))
}

func calculateAverage(numbers []int) float64 {
	total:=0
	for _,i:=range numbers {
		total+=i
	}
	
	return float64(total)/float64(len(numbers))
}

func namedReturnType(numbers []int) (avg float64) {
	avg = calculateAverage(numbers)
	return
}

func multipleReturnValues(numbers []int) (avg float64, fail bool) {
	avg = calculateAverage(numbers)
	if avg < 70.0 {
		fail = true
	}
	return
}

func variableLengthParameters(marks ...int) float64 {	
	avg := calculateAverage(marks)
	return avg
}
