//Programming languages have a lot of different types of loops like while, do while, foreach
//but go has only one that can be used in a lot different ways. for loop

package main

import "fmt"

func main() {
	i:=1
	for i<=10 {
		fmt.Println("i = ",i)
		i+=1
	}
	
	for j:=1; j<=10; j++ {
		fmt.Println("j = ", j)
	}
	
/*	//the below loop is an infinite loop, i am writing "true" after "for" for readability, even if we just write "for" it is an infinite loop
	for true {
		i+=1
		fmt.Println("I will keep on running", i)
	}*/

	for i=0;i<10;i++ {
		if i % 2 == 0{
			fmt.Println("even", i)
		}else {
			fmt.Println("odd", i)
		}
	}

	for i=0; i<10;i++{
		switch i%2 {
			case 0: 
				fmt.Println("Even : ",i)
				fmt.Println("No break")
			default:
				fmt.Println("Odd :",i)
		}
	}
}
