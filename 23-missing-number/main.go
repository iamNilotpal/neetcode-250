package main

import (
	"fmt"
	"slices"
)

func missingNumberBruteForce(nums []int) int {
	largest := len(nums)
	for i := range largest + 1 {
		if !slices.Contains(nums, i) {
			return i
		}
	}
	return -1
}

func missingNumberBetter(nums []int) int {
	frequency := make([]int, len(nums)+1)
	for _, v := range nums {
		frequency[v] += 1
	}

	for i, v := range frequency {
		if v == 0 {
			return i
		}
	}
	return -1
}

func missingNumberOptimal1(nums []int) int {
	largest := len(nums)
	totalSum := (largest * (largest + 1)) / 2

	var sumOfArrayElements int
	for _, v := range nums {
		sumOfArrayElements += v
	}

	return totalSum - sumOfArrayElements
}

func missingNumberOptimal2(nums []int) int {
	var xor1 int
	var xor2 int

	for i, v := range nums {
		xor2 ^= v
		xor1 ^= i
	}

	xor1 ^= len(nums)
	return xor1 ^ xor2
}

func main() {
	arr := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	fmt.Printf("missingNumberBruteForce: %+v \n", missingNumberBruteForce(arr))
	fmt.Printf("missingNumberBetter: %+v \n", missingNumberBetter(arr))
	fmt.Printf("missingNumberOptimal1: %+v \n", missingNumberOptimal1(arr))
	fmt.Printf("missingNumberOptimal2: %+v \n", missingNumberOptimal2(arr))
}
