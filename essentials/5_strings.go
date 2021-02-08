//strings
package main

import (
	"fmt"
)

func main() {
	book := "The color of magic"

	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// unit8 is a byte

	//access part of string using slice , syntx is index, length
	fmt.Println(book[4:11])

	//slice (no end)
	fmt.Println(book[4:])

	//slice (no start)
	fmt.Println(book[:4])

	//use + concatenate strings
	fmt.Println("t" + book[1:])

	//Unicode
	fmt.Println("It was 1/2 price!")

	//Multiline
	poem := `
	The road goes ever on
	Down from the door where it began
	...
	`
	fmt.Println(poem)
}
