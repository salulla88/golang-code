//write a function that gets a url and returns the value of content-type response HTTP header
//the function should return an error if cannot make a GET request
package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() //Make sure we close the body

	cType := resp.Header.Get("Content-Type")
	if cType == "" { //Return error if Content-Type header not found
		return "", fmt.Errorf("cant find content-type header")
	}

	return cType, nil

}

func main() {
	fmt.Println("==========Start=============")

	contentType, err := contentType("https://golang.org")

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println("Able to successfully make get request : ", contentType)
	}

	fmt.Println("==========End============")
}
