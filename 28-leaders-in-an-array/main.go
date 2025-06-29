package main

import "fmt"

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

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Printf("leadersInAnArrayBrute: %+v \n", leadersInAnArrayBrute(nums))
}
