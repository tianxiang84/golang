package main

import (
	"fmt"
)

func main() {
	//mySlide = make([]int, 10, 10)
	mySlides := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myNum := 4.0

	stat, leftNum, rightNum := findInterval(mySlides, myNum)
	fmt.Println("slides:", mySlides)
	fmt.Println("num:", myNum)
	fmt.Println("result:", stat, leftNum, rightNum)
}

func findInterval(mySlides []float64, myNum float64) (stat bool, left, right float64) {

	if myNum < mySlides[0] || myNum > mySlides[len(mySlides)-1] {
		return false, 0, 0
	} else {
		leftIdx := 0
		rightIdx := len(mySlides) - 1
		for rightIdx != leftIdx+1 {
			midIdx := int((leftIdx + rightIdx) / 2.0)
			if mySlides[midIdx] >= myNum {
				rightIdx = midIdx
			} else {
				leftIdx = midIdx
			}
		}
		return true, mySlides[leftIdx], mySlides[rightIdx]
	}
}
