package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("Program argument : ", i, os.Args[i])
	}
}
