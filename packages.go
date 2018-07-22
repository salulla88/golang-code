//An important aspect of good software is code reuse.
//In today's world, most programming is done by using one or the other existing libraries.
//Go provides a mechanism for code reuse in the form of packages.
//Go comes with a lot of inbuilt packages e.g. fmt, string, net, os, etc. Also you can create your own packages.
//In this example, we will create our own package and use it.

package main

import "fmt"
import "custom/mypackage" //this package is created in custom/mypackage directory. Once done, you install it using "go install".

func main() {
	xs := []float64{1,2,3,4}
	avg := mypackage.CalculateAverage(xs)
	fmt.Println(avg)
}
