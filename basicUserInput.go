//Scanf is the function in fmt package that can be used to get user input

package main

import "fmt"

func main() {
	fmt.Println("Please enter your name : ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Welcome", name)
}
