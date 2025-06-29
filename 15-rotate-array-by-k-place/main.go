package main

import (
	"fmt"
	"slices"
)

func rotateBetter(nums []int, k int) {
	rotations := k % len(nums)
	if rotations == 0 {
		return
	}

	temp := make([]int, rotations)
	copy(temp, nums)

	for i := rotations; i < len(nums); i++ {
		nums[i-rotations] = nums[i]
	}

	si := len(nums) - rotations
	for i := si; i < len(nums); i++ {
		nums[i] = temp[i-si]
	}
}

func rotateOptimal(nums []int, k int) {
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
	slices.Reverse(nums)
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	rotateBetter(array, 3)
	fmt.Printf("rotateBetter: %+v \n", array)

	array = []int{1, 2, 3, 4, 5, 6, 7}

	rotateOptimal(array, 3)
	fmt.Printf("rotateOptimal: %+v \n", array)
}
