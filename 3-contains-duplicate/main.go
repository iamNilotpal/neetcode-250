package main

func containsDuplicate(nums []int) bool {
	frequency := make(map[int]int)

	for _, item := range nums {
		data := frequency[item]
		if data > 0 {
			return true
		}
		frequency[item] = data + 1
	}
	return false
}

func main() {
	println("Contains Duplicate", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
