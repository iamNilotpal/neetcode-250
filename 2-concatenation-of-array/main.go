package main

import "fmt"

func getConcatenation(nums []int) []int {
	// If we're concatenating 'nums' with itself, the resulting array
	// will be exactly twice the length of the original 'nums' array.
	length := len(nums)
	ans := make([]int, 2*length)

	for i := range length {
		// Place the current element 'nums[i]' into the first half of 'ans'.
		// This covers the initial copy of 'nums'.
		ans[i] = nums[i]

		// Place the same current element 'nums[i]' into the second half of 'ans'.
		// By adding 'length' to 'i', we effectively shift the index to the
		// corresponding position in the latter half of the 'ans' array.
		ans[i+length] = nums[i]
	}

	return ans
}

func main() {
	fmt.Printf("Concatenation: %+v\n", getConcatenation([]int{1, 3, 2, 1}))
}
