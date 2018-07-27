//Go is best known for distributed, network applications. When it comes to network apps, most applications are built on top of TCP protocol in one or other way.
//Go provides with the "net" package to create a TCP based application.
//To create a tcp server is fairly simple, we use the net.Listen method

package main

import ("fmt";
	 "net";
	 "encoding/gob"
	)

func main() {
	//listen on some port
	//net.Listen accepts network type and port as arguments, here, tcp and 9000 e.g.
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err)
		return
	}

	//once you start listening, second thing is to wait for incoming connections infintely and accept them as they arrive
	for {
		con, err := ln.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}

		go handleServerConnection(con) //once connection is accepted, handle that in a goroutine
	}
}

func handleServerConnection(c net.Conn) {
	fmt.Println("handleServerConnection Started ")
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}
