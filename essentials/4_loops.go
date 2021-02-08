//Loops only for loop
package main

import (
	"fmt"
)

func main() {
	fmt.Println("========Start============")

	for i := 0; i < 3; i++ {
		fmt.Println("i : ", i)
	}

	fmt.Println("--------------")

	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println("i : ", i)
	}

	fmt.Println("--------------")
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println("i : ", i)
	}

	fmt.Println("--------------")
	i := 0
	for i < 3 {
		if i == 1 {
			break
		}
		i++
		fmt.Println("i : ", i)
	}

	fmt.Println("--------------")
	i = 0
	for {
		fmt.Println("i : ", i)
		if i == 10 {
			break
		}
		i++
		fmt.Println("i : ", i)
	}

}
