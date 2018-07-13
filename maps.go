//Maps are unordered collections of key-value pairs. Similar to associative arrays(JS), Hashmap(Java), Hashtable, and dictionaries
//Declaration : var x map[key-data-type]value-data-type; e.g. var x map[int]int

package main

import "fmt"

func main() {
	var a map[string]int
	a = make(map[string]int) //unlike arrays, maps need to be initialized unlike arrays. Try commenting this line and see what happens.
	a["Sandeep"] = 10
	a["Ravi"] = 20
	a["Ajay"] = 30

	fmt.Println("a : ", a)

	//short-hand declaration
	b:=make(map[int]int)
	for i:=0;i<10;i++ {
		b[i]=i
	}

	fmt.Println("b : ", b) //check output. You will see that the elements are not stored in the way they are inserted. i.e. unordered behaviour

	//removing elements
	delete(b, 0)
	fmt.Println("After removing elements : ", b)
	
	//accessing elements
	//see output for below statements. For missing elements, it gives default value of the data type. e.g integer 0, string ""
	fmt.Println("Element with key Sandeep and Deepak : ", a["Sandeep"], a["Deepak"])
	fmt.Println("Elements with key 1 and 10 : ", b[1], b[10])
} 
