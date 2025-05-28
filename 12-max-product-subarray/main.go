package main

func maxProductWorst(nums []int) int {
	maxProduct := 0

	for i := range nums {
		currentProduct := 1
		for j := i; j < len(nums); j++ {
			currentProduct *= nums[j]
			maxProduct = max(maxProduct, currentProduct)
		}
	}

	return maxProduct
}

func maxProductOptimal(nums []int) int {
	maxProduct := 0
	prefixMaxProduct := 1
	suffixMaxProduct := 1

	for i, v := range nums {
		if prefixMaxProduct == 0 {
			prefixMaxProduct = 1
		}
		if suffixMaxProduct == 0 {
			suffixMaxProduct = 1
		}

		prefixMaxProduct *= v
		suffixMaxProduct *= nums[len(nums)-i-1]
		maxProduct = max(maxProduct, prefixMaxProduct, suffixMaxProduct)
	}

	return maxProduct
}

func main() {
	println("maxProductWorst", maxProductWorst([]int{2, -1, -1, 3, -2, 4}))
	println("maxProductOptimal", maxProductOptimal([]int{2, -1, -1, 3, -2, 4}))
}
