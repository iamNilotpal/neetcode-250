package main

func quickSort(nums []int, si, ei int) {
	if len(nums) <= 1 || si >= ei {
		return
	}

	pivotIndex := partition(nums, si, ei)
	quickSort(nums, si, pivotIndex-1)
	quickSort(nums, pivotIndex+1, ei)
}

func partition(nums []int, si, ei int) int {
	pivot := nums[si]
	var smallOrEqualsCount int

	for i := si + 1; i <= ei; i++ {
		if nums[i] <= pivot {
			smallOrEqualsCount++
		}
	}

	pivotIndex := si + smallOrEqualsCount
	nums[pivotIndex], nums[si] = nums[si], nums[pivotIndex]

	i := si
	j := ei

	for i < pivotIndex && j > pivotIndex {
		if nums[i] <= nums[pivotIndex] {
			i++
		} else if nums[j] > nums[pivotIndex] {
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	return pivotIndex
}
