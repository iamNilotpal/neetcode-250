package main

import (
	"fmt"
	"maps"
	"slices"
)

func unionOfTwoSortedArraysBruteForce(arr1, arr2 []int) []int {
	resultSet := make(map[int]int, len(arr1)+len(arr2))

	for _, v := range arr1 {
		resultSet[v] = v
	}

	for _, v := range arr2 {
		resultSet[v] = v
	}

	sortedItems := slices.Sorted(maps.Keys(resultSet))
	return sortedItems[:]
}

func unionOfTwoSortedArraysOptimal(arr1, arr2 []int) []int {
	var i int
	var j int
	result := make([]int, 0, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		length := len(result)

		if arr1[i] == arr2[j] {
			if length <= 1 || result[length-1] != result[length-2] {
				result = append(result, arr1[i])
			}
			i++
			j++
		} else if arr1[i] < arr2[j] {
			if length <= 1 || result[length-1] != result[length-2] {
				result = append(result, arr1[i])
			}
			i++
		} else {
			if length <= 1 || result[length-1] != result[length-2] {
				result = append(result, arr2[j])
			}
			j++
		}
	}

	for i < len(arr1) {
		length := len(result)
		if length <= 1 || result[length-1] != result[length-2] {
			result = append(result, arr1[i])
		}
		i++
	}

	for j < len(arr2) {
		length := len(result)
		if length <= 1 || result[length-1] != result[length-2] {
			result = append(result, arr2[j])
		}
		j++
	}

	return result[:]
}

func main() {
	arr1 := []int{1, 2, 6, 7, 8, 9}
	arr2 := []int{2, 3, 4, 5, 10, 11, 12}

	fmt.Printf("unionOfTwoSortedArraysBruteForce: %+v \n", unionOfTwoSortedArraysBruteForce(arr1, arr2))
	fmt.Printf("unionOfTwoSortedArraysOptimal: %+v \n", unionOfTwoSortedArraysOptimal(arr1, arr2))
}
