//find the maximal value
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}

	max := 0
	for i := range nums {
		if nums[i] >= max {
			max = nums[i]
		}
	}
	fmt.Println(max)
}
