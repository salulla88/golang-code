//slices consider arrays
package main

import (
	"fmt"
)

func main() {
	fmt.Println("==========Start==============")

	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	//Length
	fmt.Println(len(loons))

	//Index
	fmt.Println(loons[1])

	//slices
	fmt.Println(loons[1:])

	//for loop
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	//single value range
	for i := range loons {
		fmt.Println(loons[i])
	}

	//double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	//double value range ignoring index by using _
	for _, name := range loons {
		fmt.Println(name)
	}

	//add elements to slice
	loons = append(loons, "elmer")
	fmt.Println(loons)
	fmt.Println("==========End==============")
}
