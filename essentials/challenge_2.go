//check if the number is even ended like 1,121,1221 likewise
package main

import (
	"fmt"
)

func main() {
	fmt.Println("========Start=========")

	count := 0

	//for every pair of 4 digit numbers
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ { //dont count twice
			n := a * b

			//if a*b  is even ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}

	fmt.Println(count)
	fmt.Println("========End=========")
}
