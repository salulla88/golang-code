//Numbers :
//Integer : uint8(byte), unit16, unint32, unint64, int8, int16, int32, int64 store 8,16,32 and 64 bits respectively.
//Machine dependent integer types : uint, int, uintptr store bits depending on the type of machine architecture in use.
//Floating-point : float32, float64 represent 32 and 64 bit numbers. complex64 and complex128 for complex numbers. In addition, they also store (0/0 = NaN) and +/-infinity
//Operations supported by numbers : +,-,*,/,%

//Strings :
//Strings are made of individual bytes, usually one for each character. 

//Booleans :
//Special 1-bit integer type used to represent 0 or 1.
//Operators supported : &&, ||, !

package main

import "fmt"

func main() {
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("9 * 8 = ", 9 * 8)
	fmt.Println("1.1 + 2.2 = ", 1.1 + 2.2)
	fmt.Println("This is a String")
	fmt.Println("True && False", true && false)
	fmt.Println("True || False", true || false)
}
