package main

import "fmt"

func isSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func main() {
	array := []int{1, 2, 3, 4, 6, 5, 6, 7, 8, 9}
	fmt.Printf("secondLargest: %+v \n", isSorted(array))
}
