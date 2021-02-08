//There are only two conditions in go. if an switch
package main

import (
	"fmt"
)

func main() {
	x := 2

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	}

	switch {
	case x > 5:
		fmt.Println("switch x greater than 5")
	default:
		fmt.Println("default case")
	}
}
