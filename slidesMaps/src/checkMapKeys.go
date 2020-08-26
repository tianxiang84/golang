package main

import (
	"fmt"
)

func main() {
	// Create a slides of keyword strings and print it
	var mySlides = make([]string, 5, 5)
	mySlides[0] = "Guangzhou"
	mySlides[1] = "Beijing"
	mySlides[2] = "Philadelphia"
	mySlides[3] = "Boston"
	mySlides[4] = "NYC"
	fmt.Println(mySlides)
	fmt.Println(cap(mySlides))

	// Create a map and print it
	var myMap = make(map[string]int)
	myMap["Guangzhou"]=10
	myMap["XinZhu"]=1
	fmt.Println(myMap)

	// Loop over the slides, if the keyword is not in the map, add it.
	for _,value := range mySlides {
	    if _, existed := myMap[value]; !existed {
	    	myMap[value] = 0
	    }
	}

	// Print the updated map
	fmt.Println(myMap)
	a,b := myMap["makeSense"]
	fmt.Println(a,b)
}