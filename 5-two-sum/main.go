package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	frequency := make(map[int]int, len(nums))
	for i, v := range nums {
		frequency[v] = i
	}

	for i, v := range nums {
		diff := target - v
		if index, ok := frequency[diff]; ok && index != i {
			return []int{i, index}
		}
	}

	return []int{}
}

func main() {
	fmt.Printf("Two Sum: %+v\n", twoSum([]int{2, 7, 11, 15}, 9))
}
