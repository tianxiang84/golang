package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(os.Args[0])
	fmt.Println(strings.Title(os.Args[0]))
   
    intYear,_ := strconv.Atoi(os.Args[1])

	fmt.Println(checkLeapYear(intYear))

	/*var inputString = "Hello"
	revertString(&inputString)
	*/
}

func checkLeapYear(year int) bool {
	if year%4 == 0 {
		return true
	} else {
		return false
	}
	//return true
}

func revertString(s1 *string) {
	//var returnString *string

    fmt.Println(*s1)
	for i:=0; i<len(*s1); i++ {
		fmt.Println(*s1)
	//	(*returnString) = (*s1)
	}

	//return *returnString
}