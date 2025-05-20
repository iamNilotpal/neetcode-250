package main

func maxSubArray(nums []int) int {
	var maxSum int = nums[0]
	var currentSum int = nums[0]

	for i := 1; i < len(nums); i++ {
		item := nums[i]
		currentSum = max(item, item+currentSum)
		maxSum = max(currentSum, maxSum)
	}

	return maxSum
}

func main() {
	println("Maximum Subarray", maxSubArray([]int{1, 2, 3, 4, -12, 6, 7, 8}))
}
