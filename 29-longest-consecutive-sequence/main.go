package main

import (
	"math"
	"slices"
)

func longestConsecutiveBruteForce(nums []int) int {
	longest := 0

	for i := range nums {
		count := 1
		x := nums[i]

		for j := 0; j < len(nums); j++ {
			if i != j && x+1 == nums[j] {
				j = 0
				x += 1
				count++
			}
		}
		longest = max(longest, count)
	}

	return longest
}

func longestConsecutiveBetter(nums []int) int {
	slices.Sort(nums)

	longest := 0
	lastLongest := 0
	lastSmallest := math.MinInt

	for i := range nums {
		if nums[i]-1 == lastSmallest {
			lastLongest++
			lastSmallest = nums[i]
		} else if nums[i] != lastSmallest {
			lastLongest = 1
			lastSmallest = nums[i]
		}
		longest = max(longest, lastLongest)
	}

	return longest
}

func longestConsecutiveOptimal(nums []int) int {
	longest := 0
	frequency := make(map[int]bool, len(nums))
	for _, v := range nums {
		frequency[v] = true
	}

	for key := range frequency {
		previousValueExists, ok := frequency[key-1]
		if ok || previousValueExists {
			continue
		}

		count := 1
		newKey := key + 1
		var isExists bool

		if exists, ok := frequency[newKey]; exists && ok {
			isExists = true
		}

		for isExists {
			count++
			newKey++
			exists, ok := frequency[newKey]
			isExists = exists && ok
		}

		longest = max(longest, count)
	}

	return longest
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	println("longestConsecutiveBruteForce", longestConsecutiveBruteForce(nums))
	println("longestConsecutiveBetter", longestConsecutiveBetter(nums))
	println("longestConsecutiveOptimal", longestConsecutiveOptimal(nums))
}
