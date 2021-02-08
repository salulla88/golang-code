//maps
package main

import (
	"fmt"
)

func main() {
	fmt.Println("===========Start============")

	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61,
	}

	//Number of items
	fmt.Println(len(stocks))

	//Get a value
	fmt.Println(stocks["MSFT"])

	//Get zero value if not foune
	fmt.Println(stocks["TSLA"])

	//Use two value to see if found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	//Set
	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	//Delete
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	//Single value for
	for key := range stocks {
		fmt.Println(key)
	}

	//Double value  for
	for key, value := range stocks {
		fmt.Println(key, ":", value)
	}

	fmt.Println("===========End============")
}
