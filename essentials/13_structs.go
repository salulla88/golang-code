//There are times when we want our data types. Go provides structs for that
//For fields in a struct, if they start with a capital letter, they are public, else they can be accessed from within the same package
package main

import "fmt"

type Trade struct {
	Symbol string  // Stock symbol
	Volume int     //Number of shares
	Price  float64 //Trade price
	Buy    bool    //true if buy trade, false if sell trade
}

func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)

	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}

	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
}
