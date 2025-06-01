package main

import (
	"fmt"
	"math"
)

func secondLargest(nums []int) int {
	largest := nums[0]
	secondLargest := math.MinInt

	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			secondLargest = largest
			largest = nums[i]
		} else if nums[i] > secondLargest && nums[i] < largest {
			secondLargest = nums[i]
		}
	}

	return secondLargest
}

func main() {
	array := []int{1, 2, 6, 3, 4, 5, 6, 7, 8, 8, 10, 9}
	fmt.Printf("secondLargest: %+v \n", secondLargest(array))
}
