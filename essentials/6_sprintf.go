package main

import (
	"fmt"
)

func main() {
	fmt.Println("==========Start===========")

	n := 42
	s := fmt.Sprintf("%d", n)

	fmt.Printf("s = %q (type of %T)\n", s, s)

	fmt.Println("==========End===========")
}
