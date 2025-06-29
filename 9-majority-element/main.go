package main

import "math"

func majorityElementWorst(nums []int) int {
	length := len(nums)

	for i, num := range nums {
		count := 1
		for j := i + 1; j < length; j++ {
			if num == nums[j] {
				count++
			}
		}
		if count > (length / 2) {
			return num
		}
	}

	return -1
}

func majorityElementBetter(nums []int) int {
	frequency := make(map[int]int, len(nums))
	for _, v := range nums {
		frequency[v] = frequency[v] + 1
	}

	for key, v := range frequency {
		if v > (len(nums) / 2) {
			return key
		}
	}

	return -1
}

func majorityElementOptimal(nums []int) int {
	count := 0
	majorityElem := math.MinInt

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count = 1
			majorityElem = nums[i]
			continue
		}

		if majorityElem == nums[i] {
			count++
		} else {
			count--
		}
	}

	return majorityElem
}

func main() {
	println("majorityElementWorst", majorityElementWorst([]int{2, 2, 1, 1, 1, 2, 2}))
	println("majorityElementBetter", majorityElementBetter([]int{2, 2, 1, 1, 1, 2, 2}))
	println("majorityElementOptimal", majorityElementOptimal([]int{2, 2, 1, 1, 1, 2, 2}))
}
