//When it comes to multithreaded programs, in languages like C and java, we have the thread creation and executor frameworks
//With Go, this has been simplified by the concepts of Goroutines.
//Whenever you want a function to be able to run alongwith another function, we use the keyword "go" followed by the function name

package main

import "fmt"


//main is the implicit goroutine. If it exits, program exits.
func main() {
	for i:=0;i<10;i++ {
		go myfunc(i) //this will call myfunc() as a goroutine
	}
	var input string
	fmt.Scanln(&input) //waiting on user input so the goroutine finishes and program does not exit as this is the last line of the program.
	
}

func myfunc(n int) {
	for i:=0; i<100; i++{
		fmt.Println(n,"I want to be called from the goroutine : ", i)
	}	
}
 
