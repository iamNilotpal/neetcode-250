package main

import "fmt"

func moveZeroesWithoutOrdering(nums []int) {
	// Intuition: Two-pointer approach to move all zeros to the end, without preserving the relative order of non-zero elements.
	// We use two pointers, `si` (start index) and `ei` (end index), that move towards each other.
	// The goal is to swap a zero at `si` with a non-zero at `ei`.

	si := 0             // 'si' starts at the beginning, looking for zeros.
	ei := len(nums) - 1 // 'ei' starts at the end, looking for non-zeros.

	// The loop continues as long as the start pointer is before the end pointer.
	for si < ei {
		if nums[si] != 0 {
			// If the element at 'si' is non-zero, it's already in its correct "left" part of the array.
			// So, move 'si' forward.
			si++
		} else if nums[ei] == 0 {
			// If the element at 'ei' is a zero, it's already in its correct "right" (zeroes) part of the array.
			// So, move 'ei' backward.
			ei--
		} else {
			// This is the crucial case: nums[si] is 0 and nums[ei] is non-zero.
			// We swap them to move the non-zero to the left and the zero to the right.
			nums[si], nums[ei] = nums[ei], nums[si]
			si++ // Move 'si' forward after the swap.
			ei-- // Move 'ei' backward after the swap.
		}
	}
}

func moveZeroesOptimalWithOrder(nums []int) {
	// Intuition: Two-pointer approach (or "snowball" technique) to move all zeros to the end,
	// while preserving the relative order of the non-zero elements.
	// This is achieved by maintaining a pointer `nextInsertIdx` that tracks where the next
	// non-zero element should be placed.

	var nextInsertIdx int // `nextInsertIdx` points to the next position where a non-zero element should be written.
	// Initially, it's 0, meaning the first non-zero element will go to index 0.

	// Iterate through the array with a 'read' pointer `i`.
	// This pointer scans for non-zero elements.
	for i := range nums {
		// If the current element `nums[i]` is not zero:
		if nums[i] != 0 {
			// This non-zero element needs to be moved to the front part of the array,
			// at the position indicated by `nextInsertIdx`.
			// We swap `nums[i]` with `nums[nextInsertIdx]`.
			// - If `i` and `nextInsertIdx` are the same, it's a swap with itself (no-op but harmless).
			// - If `i` is greater than `nextInsertIdx`, it means `nums[nextInsertIdx]` was a zero
			//   that is now being "pushed" to the right by the non-zero `nums[i]`.
			nums[nextInsertIdx], nums[i] = nums[i], nums[nextInsertIdx]
			nextInsertIdx++ // Increment `nextInsertIdx` to point to the next spot for a non-zero element.
		}
		// If `nums[i]` is 0, we simply skip it. The 'read' pointer `i` moves on,
		// but `nextInsertIdx` stays put, effectively leaving the zero behind to be
		// eventually overwritten by a non-zero, or left at the end if no more non-zeros are found.
	}
	// After the loop, all non-zero elements will be at the beginning of the array
	// (up to `nextInsertIdx-1`), and all remaining elements from `nextInsertIdx` to the end
	// will automatically be zeros.
}

func main() {
	array := []int{0, 1, 0, 3, 12, 0, 0, 1, 2, 3, 6, 9}
	moveZeroesOptimalWithOrder(array)
	fmt.Printf("moveZeroes: %+v \n", array)

	array = []int{0, 1, 0, 3, 12, 0, 0, 1, 2, 3, 6, 9}
	moveZeroesWithoutOrdering(array)
	fmt.Printf("moveZeroesWithoutOrdering: %+v \n", array)
}
