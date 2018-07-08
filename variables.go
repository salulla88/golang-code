//Variables
//declaration is "var name type"  
//short declaration "name:= value" -> here compiler figures out var type based on value.
//bulk declaration var(a=1 \n b=2 \n c=3)
//constant declaration : const name type =  value
package main

import "fmt"

func main() {
	var x string 
	x = "Sandeep"
	fmt.Println("My name is :", x)
	x+= " Lulla"
	fmt.Println("My name is :", x)
	
	var y string = "SandeepLulla"
	fmt.Println("x == y", x==y)
	x="SandeepLulla"
	fmt.Println("And now x == y : ", x==y)

	z:="Sandeep Lulla"
	fmt.Println("My name is : ", z)
	/*z = 1
	fmt.Println("Now z = ", z)*/ //This will compilation error since type cannot be changed at runtime
	
	var (a=1; b="Hi"; c=10.2) //Just declare this and only print a. It will give compilation error. If you dont use a variable dont declare it.
	fmt.Println("a :", a)
	fmt.Println("b :", b)
	fmt.Println("c :", c)

	const constVar = "I am a constant"
	fmt.Println("constVar : ", constVar)
	//constVar = "Changing const" //this will give compilation error, uncomment to see
}
