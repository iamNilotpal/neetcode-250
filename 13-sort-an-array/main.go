package main

import "fmt"

func sortArrayUsingMergeSort(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func sortArrayUsingQuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func main() {
	array := []int{5, 1, 1, 2, 0, 0}
	fmt.Printf("Before Sorting : %+v \n", array)
	array = sortArrayUsingMergeSort(array)
	fmt.Printf("After Sorting Using Merge Sort : %+v \n", array)

	array = []int{5, 1, 1, 2, 0, 0}
	fmt.Printf("Before Sorting : %+v \n", array)
	array = sortArrayUsingQuickSort(array)
	fmt.Printf("After Sorting Using Quick Sort : %+v \n", array)
}
