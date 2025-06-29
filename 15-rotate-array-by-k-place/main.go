package main

import (
	"fmt"
	"slices"
)

func rotateBetter(nums []int, k int) {
	// Intuition 1: Handle excess rotations.
	// Rotating `n` elements `n` times brings the array back to its original state.
	// So, we only care about `k` modulo `len(nums)` rotations.
	// This reduces unnecessary operations for very large `k`.
	rotations := k % len(nums)

	// If rotations is 0, no rotation is needed.
	if rotations == 0 {
		return
	}

	// Intuition 2: Use a temporary buffer for the elements that move from the end to the front.
	// When rotating `nums` by `rotations`, the last `rotations` elements will move to the beginning.
	// We need to save these elements before they are overwritten.
	// `temp` will store these `rotations` elements.
	// `copy(temp, nums[len(nums)-rotations:])` would be more direct to copy the last elements.
	// The original code copies the *first* `rotations` elements, which is a different approach.
	// Let's analyze the original code's copy: `copy(temp, nums)`. This means `temp` gets `nums[0]` to `nums[rotations-1]`.
	// This suggests a strategy where the first part is stored, the rest is shifted, and then the stored part is placed.
	// This is a common strategy for rotations, but the index logic for shifting will be different.

	// Re-evaluating the original code's strategy based on `copy(temp, nums)`:
	// It appears the intent is to move the first 'rotations' elements to a temporary buffer.
	// Then, shift the remaining `len(nums) - rotations` elements to the front.
	// Finally, place the elements from `temp` at the end.
	// However, this approach would rotate *left* by `rotations` (elements from front go to back).
	// The problem "rotate array" usually implies *right* rotation (elements from back go to front).

	// Assuming the problem is standard "right rotation":
	// The initial `copy(temp, nums)` copies the *first* `rotations` elements into `temp`.
	// This doesn't seem to align with a standard right rotation strategy using a temp array
	// where the *last* `rotations` elements are moved to the temp array.
	// Let's trace the standard approach for a right rotation first, and then the given code's logic.

	// --- Standard Right Rotation using temp array (save last 'rotations' elements) ---
	//   temp := make([]int, rotations)
	//   copy(temp, nums[len(nums)-rotations:]) // Save the last 'rotations' elements
	//   // Shift the first `len(nums) - rotations` elements to the right by `rotations` positions
	//   for i := len(nums) - 1 - rotations; i >= 0; i-- {
	//       nums[i+rotations] = nums[i]
	//   }
	//   // Place the saved elements at the beginning
	//   copy(nums, temp)
	//   return
	// --- End of Standard Right Rotation ---

	// --- Analyzing the provided `rotateBetter` function's specific strategy ---
	// The provided code copies `nums[0]` to `nums[rotations-1]` into `temp`.
	temp := make([]int, rotations) // Creates a temp slice of size `rotations`.
	copy(temp, nums)               // Copies elements from `nums[0]` up to `nums[rotations-1]` into `temp`.
	// Example: nums=[1,2,3,4,5], k=2, rotations=2. temp will be [1,2]

	// Intuition 3: Shift the remaining elements.
	// This loop shifts elements from index `rotations` onwards to the left.
	// `nums[i-rotations]` means the element at `nums[i]` is moved `rotations` steps to the left.
	// Example: nums=[1,2,3,4,5], temp=[1,2], rotations=2
	// i=2: nums[0] = nums[2] (3) -> nums=[3,2,3,4,5]
	// i=3: nums[1] = nums[3] (4) -> nums=[3,4,3,4,5]
	// i=4: nums[2] = nums[4] (5) -> nums=[3,4,5,4,5]
	// After this loop, nums becomes `[3,4,5,4,5]` (conceptually `[3,4,5, _, _]`). The original first `rotations` elements (1,2) are gone.
	for i := rotations; i < len(nums); i++ {
		nums[i-rotations] = nums[i]
	}

	// Intuition 4: Place the temporarily stored elements at the end.
	// `si` is the starting index in `nums` where the `temp` elements should be placed.
	// This is the index after all the shifted elements.
	si := len(nums) - rotations
	// Iterate from `si` to the end of `nums`.
	// For each position, place an element from `temp`.
	// `i-si` correctly maps the index in `nums` to the corresponding index in `temp`.
	// Example (continuing): nums=[3,4,5,4,5], temp=[1,2], rotations=2, si=3
	// i=3: nums[3] = temp[3-3] = temp[0] (1) -> nums=[3,4,5,1,5]
	// i=4: nums[4] = temp[4-3] = temp[1] (2) -> nums=[3,4,5,1,2]
	// Final nums: [3,4,5,1,2]
	// This result ([3,4,5,1,2]) is a *left* rotation by `len(nums) - k` positions,
	// or effectively a *right* rotation by `k` positions.
	// For original nums=[1,2,3,4,5], k=2 (right rotation):
	// Expected: [4,5,1,2,3]
	// Actual: [3,4,5,1,2]
	// It seems this specific implementation results in a *left* rotation by `k` positions.
	// A left rotation by `k` is equivalent to a right rotation by `len(nums) - k`.
	// So, the name `rotateBetter` is accurate, but it's a left rotation implementation.

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
