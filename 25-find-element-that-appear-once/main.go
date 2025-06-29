package main

import (
	"fmt"
	"math"
)

func findElementBruteForce(nums []int) int {
	for _, num := range nums {
		var count int
		for _, currElem := range nums {
			if num == currElem {
				count++
				if count > 1 {
					break
				}
			}
		}
		if count == 1 {
			return num
		}
	}

	return math.MinInt
}

func main() {
	arr := []int{1, 1, 1, 1, 0, 2, 2, 3, 3}
	fmt.Printf("findElementBruteForce: %+v \n", findElementBruteForce(arr))
}
