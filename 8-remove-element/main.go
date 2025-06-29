package main

func removeElement(nums []int, val int) int {
	// Intuition: Two-Pointer Approach (In-place modification)
	// The problem asks to remove all occurrences of `val` in-place and return the new length.
	// The order of elements *not* equal to `val` can be changed.
	// This means we can effectively "shift" the elements we want to keep to the beginning of the array.

	// `nextInsertIndex` acts as a "write pointer". It tracks the position where the
	// next non-`val` element should be placed. It also ultimately represents the
	// new length of the modified array.
	var nextInsertIndex int

	// Iterate through the original array using a "read pointer" `i`.
	for i := range nums {
		// If the current element `nums[i]` is NOT the value we want to remove:
		if nums[i] != val {
			// It means this element should be kept.
			// Place `nums[i]` at the position indicated by `nextInsertIndex`.
			// Since `nextInsertIndex` starts at 0 and only increments for elements we keep,
			// it ensures all non-`val` elements are moved to the front of the array
			// in their original relative order.
			nums[nextInsertIndex] = nums[i]

			// Advance the `nextInsertIndex` to prepare for the next non-`val` element.
			nextInsertIndex++
		}
	}

	// Intuition: The `nextInsertIndex` now holds the count of elements that were NOT `val`.
	// This is the new length of the array after removal.
	return nextInsertIndex
}

func main() {
	println("Remove Element", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
