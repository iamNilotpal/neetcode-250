package main

import "fmt"

func rearrangeArrayBruteForce(nums []int) []int {
	length := len(nums)
	positives := make([]int, 0, (length / 2))
	negatives := make([]int, 0, (length / 2))

	for _, v := range nums {
		if v >= 0 {
			positives = append(positives, v)
		} else {
			negatives = append(negatives, v)
		}
	}

	var i, j, k int
	for i < len(positives) && j < len(negatives) && k < len(nums) {
		nums[k] = positives[i]
		nums[k+1] = negatives[j]
		k += 2
		i++
		j++
	}

	return nums
}

func main() {
	nums := []int{3, 1, -2, -5, 2, -4}
	fmt.Printf("rearrangeArrayBruteForce: %+v \n", rearrangeArrayBruteForce(nums))
}
