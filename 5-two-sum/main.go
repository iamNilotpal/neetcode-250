package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// Intuition 1: Store numbers and their indices for quick lookups.
	// We need to efficiently check if a "complement" (target - current_number) exists
	// in the array. A hash map (or dictionary) is perfect for this, as it allows
	// O(1) average time complexity for lookups.
	// The map will store: key = number from nums, value = its index.
	// Pre-allocating capacity can offer a minor performance improvement.
	frequency := make(map[int]int, len(nums))
	// Populate the map: Iterate through 'nums' once to store each number and its index.
	// If there are duplicate numbers, this approach stores the *last* seen index.
	// This is important for handling cases where the complement might be a duplicate
	// of the current number itself, but at a different index.
	for i, v := range nums {
		frequency[v] = i
	}

	// Intuition 2: For each number, calculate its required complement and look it up.
	// Iterate through 'nums' a second time. For each number `v` at index `i`:
	for i, v := range nums {
		// Calculate the 'diff' (complement) needed to reach the 'target'.
		// If `v` + `diff` = `target`, then `diff` = `target` - `v`.
		diff := target - v

		// Check if this 'diff' exists as a key in our map.
		// `index, ok := frequency[diff]` tries to retrieve the value (index) associated with 'diff'.
		// `ok` will be true if 'diff' is found in the map, false otherwise.
		//
		// Crucially, we also need to ensure that the `index` found for `diff` is NOT
		// the same as the current index `i`. This prevents using the same element twice
		// to form the sum (e.g., if target is 6 and current number is 3, we don't want
		// to pick 3 twice from the same index).
		if index, ok := frequency[diff]; ok && index != i {
			// If 'diff' is found and its index is different from the current index `i`,
			// we have found the two numbers that sum up to the target.
			// Return their indices as a slice.
			return []int{i, index}
		}
	}

	// Intuition 3: If no such pair is found after checking all numbers.
	// If the loop completes without finding a pair, it means no two numbers
	// in the array sum up to the target. Return an empty slice as per problem requirements.
	return []int{}
}

func main() {
	fmt.Printf("Two Sum: %+v\n", twoSum([]int{2, 7, 11, 15}, 9))
}
