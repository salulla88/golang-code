//defer is a garbage collector in go. it can be used to clean up any resource. it will be called even if exception occurs
package main

import "fmt"

func cleanup(name string) {
	fmt.Printf("CLeaning up %s\n", name)
}

func main() {
	defer cleanup("A")
	defer cleanup("B") //defer calls are made in reverse order
	fmt.Println("worker")
}
