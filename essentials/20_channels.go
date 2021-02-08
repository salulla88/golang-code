package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("========Start=========")
	ch := make(chan int)

	/* this will create a block as we havent pushed anything and are trying to receive
	<-ch
	fmt.Println("Here")
	*/

	go func() {
		//Send number of the channel
		ch <- 353
	}()

	val := <-ch
	fmt.Printf("got %d\n", val)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending : %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received : %d\n", i)
	}

}
