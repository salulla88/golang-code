//panic and recover. panic is like exception
package main

import "fmt"

func main() {
	//part 1
	// vals := []int{1, 2, 3}
	// //This will cause a panic
	// v := vals[10]
	// fmt.Println(v)

	//part 2
	// file, err := os.Open("no-such-file")
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	// fmt.Println("file opened")

	//part 3
	v := safeValue([]int{1, 2, 3}, 0)
	fmt.Println(v)
}

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR %s\n", err)
		}
	}()
	return vals[index]
}
