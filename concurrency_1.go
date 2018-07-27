//How to do interthread communication? i.e. how to communicate between 2 goroutines.
//Go provides "channels" for that purpose

package main

import "fmt"

//channels have both sending & receiving capabilities, you can restrict the functions from doing the same. "<- {channame}" means receiving while "{channame} <-" means sending capacity.
//if we try to receive in a function that has channel with sending capability only, then it will be caught at compilation only. Thanks to Go :) You dont have to wait for runtime exceptions
//channels can take any type of variables, lets use string for sample here.

func receiveMessage(rc chan string) {
	for {
		fmt.Println("Received : ", <-rc)
	}
}

func sendMessage(sc chan string) {
	for {
		sc <- "Heart beat"
	}
}

func main() {
	var channel chan string = make(chan string)

	go receiveMessage(channel)
	go sendMessage(channel)

	var input string
	fmt.Scanln(&input)
}
