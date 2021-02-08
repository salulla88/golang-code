//go routines
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	fmt.Printf(time.Now().String())
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {

	fmt.Println("========Start=========")

	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	//sync
	for _, url := range urls {
		go returnType(url)
	}

	var wg sync.WaitGroup //used to wait for all routines to complete and then show main
	for _, url := range urls {
		wg.Add(1)
		//async
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()

	fmt.Println("========End=========")

}
