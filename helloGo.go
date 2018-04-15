//This is the way you write comments in Go.
//Go code is organized into packages. A package consists of one or more .go files.
//Each go source file begins with the package declaration, that states which package
//it belongs to, here package 'main', followed by the packages that it imports.
//A go program does not compile if there are missing imports or unnecessary imports.
//This prevents references to unncessary packages.

package main //main package is special, which tells that this is a executable program and not a library.

import "fmt" //fmt is a go package for printing formatted output and scanning input.

//function main is also a special declaration. It tells the compiler that this is where the execution of the program will begin.
func main() {
	fmt.Println("Hello Go!!")
}
