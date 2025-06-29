package main

import "slices"

func removeElement(nums []int, val int) int {
	var nextInsertIndex int

	for i := range nums {
		if nums[i] != val {
			nums[nextInsertIndex] = nums[i]
			nextInsertIndex++
		}
	}

	nums = slices.Delete(nums, nextInsertIndex, len(nums))
	return len(nums)
}

func main() {
	println("Remove Element", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
