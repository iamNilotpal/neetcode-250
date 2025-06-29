package main

// majorityElementWorst: Brute-force approach with O(n^2) time complexity.
func majorityElementWorst(nums []int) int {
	// Intuition: For each number in the array, count its occurrences
	// by comparing it with every other number. If any count exceeds n/2,
	// that's the majority element.

	length := len(nums)

	// Outer loop: Pick each number 'num' to check if it's the majority element.
	for i, num := range nums {
		count := 1 // Initialize count for the current 'num'. Start with 1 because 'num' itself counts.

		// Inner loop: Compare 'num' with all subsequent elements to find duplicates.
		// Start from 'i + 1' to avoid re-counting or comparing with itself at the same index.
		for j := i + 1; j < length; j++ {
			if num == nums[j] {
				count++ // Increment count if a duplicate is found.
			}
		}

		// After checking all other elements, if 'count' is greater than half the array's length,
		// 'num' is the majority element.
		if count > (length / 2) {
			return num
		}
	}

	// This line is typically unreachable in problems where a majority element is guaranteed to exist.
	// It's a safeguard for cases where it might not exist or for an empty input.
	return -1
}

// majorityElementBetter: Using a Hash Map (Frequency Map) with O(n) time complexity and O(n) space complexity.
func majorityElementBetter(nums []int) int {
	// Intuition: Store the frequency of each number.
	// The majority element will be the one with a frequency greater than n/2.

	// 1. Create a frequency map:
	// 'frequency' will store numbers as keys and their counts as values.
	// Pre-allocate capacity for potential minor performance gain.
	frequency := make(map[int]int, len(nums))

	// 2. Populate the frequency map:
	// Iterate through 'nums' once, incrementing the count for each number.
	for _, v := range nums {
		frequency[v] = frequency[v] + 1
	}

	// 3. Find the majority element from the map:
	// Iterate through the key-value pairs in the frequency map.
	for key, v := range frequency {
		// If the count 'v' for any 'key' is greater than half the array's length,
		// then 'key' is the majority element.
		if v > (len(nums) / 2) {
			return key
		}
	}

	// This line is typically unreachable if a majority element is guaranteed.
	return -1
}

// majorityElementOptimal: Boyer-Moore Majority Vote Algorithm with O(n) time complexity and O(1) space complexity.
func majorityElementOptimal(nums []int) int {
	// Intuition: If we have a majority element, it will outvote all other elements combined.
	// This algorithm leverages a "vote" system:
	// - When we see the current `majorityElem`, we increment its `count`.
	// - When we see a different element, we decrement `count`.
	// - If `count` drops to 0, it means the current `majorityElem` has been "cancelled out"
	//   by non-majority elements. We then pick the next element as a new potential `majorityElem`.
	// Because the true majority element appears more than n/2 times, it will always be the one
	// remaining with a positive count at the end.

	count := 0              // Stores the current vote count for `majorityElem`.
	majorityElem := nums[0] // Initialize `majorityElem` with the first element.

	for i := range nums {
		if count == 0 {
			// If count is 0, it means the previous potential majority element has been "canceled out".
			// Assume the current element `nums[i]` is the new potential majority element.
			majorityElem = nums[i]
			count = 1 // Start its vote count at 1.
		} else if majorityElem == nums[i] {
			// If the current element matches the `majorityElem`, increment its vote count.
			count++
		} else {
			// If the current element does not match the `majorityElem`, decrement its vote count.
			// This effectively cancels out one vote for the `majorityElem`.
			count--
		}
	}

	// After iterating through the entire array, `majorityElem` will hold the majority element.
	// (This algorithm guarantees correctness IF a majority element exists).
	return majorityElem
}

func main() {
	println("majorityElementWorst", majorityElementWorst([]int{2, 2, 1, 1, 1, 2, 2}))
	println("majorityElementBetter", majorityElementBetter([]int{2, 2, 1, 1, 1, 2, 2}))
	println("majorityElementOptimal", majorityElementOptimal([]int{2, 2, 1, 1, 1, 2, 2}))
}
