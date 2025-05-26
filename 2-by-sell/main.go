package main

func maxProfit(prices []int) int {
	maxProfit := 0
	minPriceSoFar := prices[0]

	for i := 1; i < len(prices); i++ {
		maxProfit = max(maxProfit, prices[i])
		minPriceSoFar = min(minPriceSoFar, prices[i])
	}

	return maxProfit
}

func main() {
	println("Max Profit", maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
