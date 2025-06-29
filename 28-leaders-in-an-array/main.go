package main

import (
	"fmt"
	"math"
)

func leadersInAnArrayBrute(nums []int) []int {
	result := make([]int, 0)

	for i := range nums {
		isLeader := true

		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				isLeader = false
			}
		}

		if isLeader {
			result = append(result, nums[i])
		}
	}

	return result
}

func leadersInAnArrayOptimal(nums []int) []int {
	max := math.MinInt
	result := make([]int, 0)

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > max {
			max = nums[i]
			result = append(result, nums[i])
		}
	}

	return result
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Printf("leadersInAnArrayBrute: %+v \n", leadersInAnArrayBrute(nums))

	nums = []int{10, 22, 12, 3, 0, 6}
	fmt.Printf("leadersInAnArrayOptimal: %+v \n", leadersInAnArrayOptimal(nums))
}
