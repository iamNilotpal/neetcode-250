package main

import "fmt"

func getConcatenation(nums []int) []int {
	length := len(nums)
	ans := make([]int, 2*length)

	for i := range length {
		ans[i] = nums[i]
		ans[i+length] = nums[i]
	}
	return ans
}

func main() {
	fmt.Printf("Concatenation: %+v\n", getConcatenation([]int{1, 3, 2, 1}))
}
