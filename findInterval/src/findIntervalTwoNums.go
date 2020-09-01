package main

import "fmt"

func main() {
	//mySlide = make([]int, 10, 10)
	mySlides := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myNum := []float64{4.3, 5.6}

	fmt.Printf("mySlides is of type %T\n", mySlides)
	fmt.Printf("mySlides is of type %T\n", myNum)

	currL := 0
	currR := len(mySlides)
	for _, num := range myNum {
		fmt.Println(num)
		stat, leftNum, rightNum, leftIdx, _ := findInterval(mySlides[currL:currR], num)
		fmt.Println(stat, leftNum, rightNum)
		currL = leftIdx
	}
}

func findInterval(mySlides []float64, myNum float64) (stat bool, left, right float64, leftIdx, rightIdx int) {

	if myNum < mySlides[0] || myNum > mySlides[len(mySlides)-1] {
		return false, 0, 0, 0, 0
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
		return true, mySlides[leftIdx], mySlides[rightIdx], leftIdx, rightIdx
	}
}
