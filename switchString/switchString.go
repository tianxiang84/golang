package main

import (
	"fmt"
)

func main() {
	var (
		string1 = "This is string 1."
		string2 = "This is string 2."
	)

	fmt.Println("Hello","World!")

	fmt.Println(string1)
	fmt.Println(string2)

	switchString(&string1, &string2)

	fmt.Println(string1)
	fmt.Println(string2)
}

func switchString(string1,string2 *string){
	temp := *string2
	*string2 = *string1
	*string1 = temp
}