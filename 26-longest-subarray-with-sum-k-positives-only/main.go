package main

func subarraySumBruteForce(nums []int, k int) int {
	var maxCount int

	for i := range nums {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				maxCount = max(maxCount, (j - i + 1))
				break
			}
		}
	}

	return maxCount
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 1, 3, 1, 1, 2, 3, 4, 5}
	println("subarraySumBruteForce:", subarraySumBruteForce(nums, 9))
}
