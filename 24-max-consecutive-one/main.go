package main

import "fmt"

func findMaxConsecutiveOnesBruteForce(nums []int) int {
	var i int
	var maxCount int

	for i < len(nums) {
		if nums[i] != 1 {
			i++
			continue
		}

		j := i + 1
		count := 1

		for j < len(nums) {
			if nums[j] == 1 {
				j++
				count++
			} else {
				i = j + 1
				break
			}
		}

		maxCount = max(maxCount, count)
		if j >= len(nums) {
			break
		}
	}

	return maxCount
}

func findMaxConsecutiveOptimal(nums []int) int {
	var count int
	var maxCount int

	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			count = 0
		}
		maxCount = max(maxCount, count)
	}

	return maxCount
}

func main() {
	arr := []int{1, 0, 1, 1, 0, 1, 1, 1}
	fmt.Printf("findMaxConsecutiveOnesBruteForce: %+v \n", findMaxConsecutiveOnesBruteForce(arr))
	fmt.Printf("findMaxConsecutiveOptimal: %+v \n", findMaxConsecutiveOptimal(arr))
}
