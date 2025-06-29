package main

func maxSumSubArrayWorst(nums []int) int {
	var maxSum int

	for i := range nums {
		var currentSum int
		for j := i; j < len(nums); j++ {
			currentSum += nums[j]
			maxSum = max(maxSum, currentSum)
		}
	}
	return maxSum
}

func maxSumSubArrayOptimal(nums []int) int {
	var maxSum int
	var currentSum int

	for _, v := range nums {
		if currentSum < 0 {
			currentSum = 0
		}
		currentSum += v
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func main() {
	println("maxSumSubArrayWorst", maxSumSubArrayWorst([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	println("maxSumSubArrayOptimal", maxSumSubArrayOptimal([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
