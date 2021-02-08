package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
{
	"user":"Sandeep",
	"type":"deposit",
	"amount": 100000.3
}
`

//Request is a bank transaction
type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	fmt.Println("====Start=======")
	rdr := bytes.NewBufferString(data) //Simulate a file/socket
	//Decode request
	dec := json.NewDecoder(rdr)

	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("error: cant decode - %s", err)
	}

	fmt.Printf("got: %+v\n", req)

	//Create response
	prevBalance := 8500000.0 //Loaded from database
	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	//Encode response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: cant encode - %s", err)
	}

	fmt.Println("====End=======")
}
