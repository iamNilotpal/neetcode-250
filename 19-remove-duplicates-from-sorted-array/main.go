package main

import "fmt"

func removeDuplicates(nums []int) int {
	nextInsertIdx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[nextInsertIdx] = nums[i]
			nextInsertIdx++
		}
	}

	fmt.Printf("NUMS: %+v \n", nums)
	return nextInsertIdx
}

func main() {
	array := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("secondLargest: %+v \n", removeDuplicates(array))
}
