//select can be used for polling from channels
//one of the most useful cases for select is timeout that is demonstrated below
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=======Start========")

	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 42
	}()

	select {
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("got %d from ch2\n", val)
	}

	fmt.Println("-------------------")

	out := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f\n", val)
	case <-time.After(20 * time.Millisecond):
		fmt.Printf("timeout")
	}

	fmt.Println("=======End========")
}
