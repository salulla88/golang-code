//Go provides the "strings" package for string manipulation.

package main

import "fmt"
import "strings"

func main() {
	str:="This is a rainy day"
	search:="rainy"

	//contains
	fmt.Println(strings.Contains(str, search))
	fmt.Println(strings.Contains("This is a rainy day", "rainy"))
	
	//starts with
	fmt.Println(strings.HasPrefix("Thanksgiving", "Thank")) 
	fmt.Println(strings.HasPrefix("Thanksgiving", "thank"))//notice its case-sensitive

	//ends with
	fmt.Println(strings.HasSuffix("Thanksgiving", "ing")) 

	//count occurences
	fmt.Println(strings.Count("SandeepSandeepSandeep", "sandeep"))
	fmt.Println(strings.Count("SandeepSandeepSandeep", "Sandeep"))

	//index of
	fmt.Println(strings.Index("Sandeep", "l"))
	fmt.Println(strings.Index("Sandeep", "e"))

	//to lower
	fmt.Println(strings.ToLower("Sandeep"))
	fmt.Println(strings.ToLower("sandeep"))

	//to upper
	fmt.Println(strings.ToUpper("Sandeep"))
	fmt.Println(strings.ToUpper("sandeep"))

	//concate
	fmt.Println(strings.Join([]string{"San", "deep"}, ""))	

	//split
	fmt.Println(strings.Split("Today is Sunday", " "))

	//likewise there are many more functions available
}
