//Go provides mechanism for reading, writing, creating files using os, ioutils packages

package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	//create a file & get its handle
	file, err := os.Create("test.txt")
	
	//if file could not be created, then we get something in err
	if err != nil {
		fmt.Println("failed to create file.")
		return
	}

	//write to file
	file.WriteString("Hi i was created using Go.")
	
	//we have got the file handle, so we ensure that file gets closed at the end of this program
	defer file.Close()

	//read the file
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("error reading file")
		return
	}

	str := string(bs)
	fmt.Println("File Contents : ", str)

	//lets see what happens when a file does not exist and we read it

	bs, err = ioutil.ReadFile("test1.txt")
	if err != nil {
		fmt.Println("error reading file : ", err)
		return
	}
}
