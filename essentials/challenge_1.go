//if the number is divisible by 3 say "fizz", if divisible by 5, say "buzz", if by both say "fizz buzz"
package main

import (
	"fmt"
)

func main() {
	fmt.Println("===========Start============")

	i := 0
	for {
		i++
		switch {
		case (i%3 == 0 && i%5 == 0):
			fmt.Println("fizz buzz")
		case i%5 == 0:
			fmt.Println("buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		default:
			fmt.Println("--------------------")
		}

		if i == 20 {
			break
		}
	}
}
