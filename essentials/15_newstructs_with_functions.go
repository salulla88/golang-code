//constructor like create methods for structs
package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string  // Stock symbol
	Volume int     //Number of shares
	Price  float64 //Trade price
	Buy    bool    //true if buy trade, false if sell trade
}

func NewTrade(s string, v int, p float64, b bool) (*Trade, error) {
	if s == "" {
		return nil, fmt.Errorf("symbol cant be empty")
	}

	if v <= 0 {
		return nil, fmt.Errorf("volume must be >=0 (was %d)", v)
	}

	if p <= 0.0 {
		return nil, fmt.Errorf("price must be >=0 (was %f)", p)
	}

	trade := &Trade{
		Symbol: s,
		Volume: v,
		Price:  p,
		Buy:    b,
	}

	return trade, nil
}

func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	fmt.Println("=========Start==========")

	t, err := NewTrade("MSFT", 100, 99.98, false)

	if err != nil {
		fmt.Printf("error: cant create trade - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.Value())

	fmt.Println("=========End==========")
}
