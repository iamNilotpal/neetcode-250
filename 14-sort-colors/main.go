package main

import "fmt"

func sortColorsTwoPass(nums []int) {
	var zerosCount, onesCount int
	for _, v := range nums {
		if v == 0 {
			zerosCount++
		} else if v == 1 {
			onesCount++
		}
	}

	for i := range zerosCount {
		nums[i] = 0
	}

	for i := zerosCount; i < zerosCount+onesCount; i++ {
		nums[i] = 1
	}

	for i := zerosCount + onesCount; i < len(nums); i++ {
		nums[i] = 2
	}
}

func sortColorsOnePass(nums []int) {
	si := 0
	mid := 0
	ei := len(nums) - 1

	for mid <= ei {
		if nums[mid] == 1 {
			mid++
		} else if nums[mid] == 0 {
			nums[si], nums[mid] = nums[mid], nums[si]
			si++
			mid++
		} else {
			nums[mid], nums[ei] = nums[ei], nums[mid]
			ei--
		}
	}
}

func main() {
	input := []int{2, 0, 2, 1, 1, 0}
	fmt.Printf("Before Sorting : %+v \n", input)
	sortColorsTwoPass(input)
	fmt.Printf("After Sorting Using sortColorsTwoPass : %+v \n", input)

	input = []int{2, 0, 2, 1, 1, 0}
	fmt.Printf("Before Sorting : %+v \n", input)
	sortColorsOnePass(input)
	fmt.Printf("After Sorting Using sortColorsOnePass : %+v \n", input)
}
