package main

func mergeSort(nums []int, si, ei int) {
	if si >= ei {
		return
	}

	mid := (si + ei) / 2
	mergeSort(nums, si, mid)
	mergeSort(nums, mid+1, ei)
	merge(nums, si, ei)
}

func merge(nums []int, si, ei int) {
	mid := (si + ei) / 2
	length := (ei - si) + 1
	result := make([]int, length)

	k := 0
	i := si
	j := mid + 1

	for i <= mid && j <= ei {
		if nums[i] < nums[j] {
			result[k] = nums[i]
			i++
		} else {
			result[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		result[k] = nums[i]
		i++
		k++
	}

	for j <= ei {
		result[k] = nums[j]
		j++
		k++
	}

	k = 0
	for i := si; i <= ei; i++ {
		nums[i] = result[k]
		k++
	}
}
