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
	prefix := 1
	suffix := 1
	maxProduct := 0

	for i, v := range nums {
		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}

		prefix *= v
		suffix *= nums[len(nums)-i-1]
		maxProduct = max(maxProduct, prefix, suffix)
	}

	return maxProduct
}

func main() {
	println("maxProductWorst", maxProductWorst([]int{2, -1, -1, 3, -2, 4}))
	println("maxProductOptimal", maxProductOptimal([]int{2, -1, -1, 3, -2, 4}))
}
