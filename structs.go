//Structs are a way to define your own data types in Go.
//Declaration type <name> struct { field datatype, ...}

package main

import "fmt"

type Employee struct {
	id int
	name string
	age int
	role string
}

func main() {
	//first way to create struct instance
	var a Employee
	a.id = 1
	a.name = "Suraj"
	a.age = 35
	a.role = "HR"

	fmt.Println("a : ", a)

	//second way to create struct instance
	var b = new(Employee)
	fmt.Println("Before assigning values b : ", b)
	b.id = 2
	b.name = "Sunil"
	b.age = 27
	b.role = "Admin"

	fmt.Println("After assigning values b : ", b)

	//third way : this one also returns a pointer as in case of second way
	c:= &Employee{3, "Ravi", 20, "Intern"}
	fmt.Println("c : ", *c)
	fmt.Println("Using struct method : ")
	c.getRole()
}

//Structs can have methods specific to them. Just write pointer to struct after "func" and before function name
func (e *Employee) getRole() {
	fmt.Println("Employee name : " + e.name + " Role : " + e.role)
}


