package main

import (
	"fmt"
)

func intersectionBruteForce(nums1 []int, nums2 []int) []int {
	minLength := min(len(nums1), len(nums2))
	frequency := make(map[int]int, minLength)

	toIterateArray := nums1
	if len(nums2) < len(nums1) {
		toIterateArray = nums2
	}

	for _, v := range toIterateArray {
		frequency[v] += 1
	}

	if len(nums1) > len(nums2) {
		toIterateArray = nums1
	} else {
		toIterateArray = nums2
	}

	result := make([]int, 0, minLength)
	for _, num := range toIterateArray {
		frequencyVal, ok := frequency[num]
		if !ok || frequencyVal == 0 {
			continue
		}

		result = append(result, num)
		frequency[num] = 0
	}

	return result[:]
}

func intersectionOptimal(nums1 []int, nums2 []int) []int {
	var i int
	var j int
	minLength := min(len(nums1), len(nums2))
	result := make([]int, 0, minLength)

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return result[:]
}

func main() {
	arr1 := []int{4, 5, 5, 6, 9}
	arr2 := []int{4, 5, 8, 9}

	fmt.Printf("intersectionBruteForce: %+v \n", intersectionBruteForce(arr1, arr2))
	fmt.Printf("intersectionOptimal: %+v \n", intersectionOptimal(arr1, arr2))
}
