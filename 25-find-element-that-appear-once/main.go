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

func findElementBetter(nums []int) int {
	frequency := make(map[int]int, len(nums))
	for _, v := range nums {
		frequency[v] += 1
	}

	for key, val := range frequency {
		if val == 1 {
			return key
		}
	}
	return math.MinInt
}

func findElementOptimal(nums []int) int {
	var xor int
	for _, v := range nums {
		xor ^= v
	}
	return xor
}

func main() {
	arr := []int{1, 1, 1, 1, 0, 2, 2, 3, 3}
	fmt.Printf("findElementBruteForce: %+v \n", findElementBruteForce(arr))
	fmt.Printf("findElementBetter: %+v \n", findElementBetter(arr))
	fmt.Printf("findElementOptimal: %+v \n", findElementOptimal(arr))
}
