//In the previous example, we created tcp server.
//Here we create a tcp  client that connects to that server and sends message

package main

import (
	"fmt";
	"net";	
	"encoding/gob"
	)

func main() {
	// connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
		return
	}

	// send the message
	msg := "Hello Server"
	err = gob.NewEncoder(conn).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	conn.Close()
}
