package main

import "fmt"

func moveZeroesWithoutOrdering(nums []int) {
	si := 0
	ei := len(nums) - 1

	for si < ei {
		if nums[si] != 0 {
			si++
		} else if nums[ei] == 0 {
			ei--
		} else {
			nums[si], nums[ei] = nums[ei], nums[si]
			si++
			ei--
		}
	}
}

func moveZeroesOptimalWithOrder(nums []int) {
	var nextInsertIdx int
	for i := range nums {
		if nums[i] != 0 {
			nums[nextInsertIdx], nums[i] = nums[i], nums[nextInsertIdx]
			nextInsertIdx++
		}
	}
}

func main() {
	array := []int{0, 1, 0, 3, 12, 0, 0, 1, 2, 3, 6, 9}
	moveZeroesOptimalWithOrder(array)
	fmt.Printf("moveZeroes: %+v \n", array)

	array = []int{0, 1, 0, 3, 12, 0, 0, 1, 2, 3, 6, 9}
	moveZeroesWithoutOrdering(array)
	fmt.Printf("moveZeroesWithoutOrdering: %+v \n", array)
}
