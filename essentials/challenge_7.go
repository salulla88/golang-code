//interfaces
package main

import (
	"fmt"
	"io"
)

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {

}

func main() {
	fmt.Println("==========Start============")

	fmt.Println("==========End============")
}
