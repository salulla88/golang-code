//Go has a special statement called "defer" that schedules a function call to be run after the function completes.
//Consider of it as a "finally" block in java that will execute no matter what happens within the try block
//There are three such statements in Go : defer, panic, recover. Panic is like throwing exception, recover is like catching exception even if it is runtime.

package main

import "fmt"

func main() {
	a1 := func() {
		fmt.Println("I am a1.")
	}

	a2 := func() {
		fmt.Println("I am a2.")
	}

	panicRecover()

   	defer a2() //see that due to defer being added the output of the program always runs a1 first.
	a1()
}

func panicRecover() {
	defer func() {
		str := recover()
		fmt.Println("Cause of panic : ", str)
	}()

	panic("I am a panic.")
} 
