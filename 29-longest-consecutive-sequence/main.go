package main

func longestConsecutiveBruteForce(nums []int) int {
	longest := 0

	for i := range nums {
		count := 1
		x := nums[i]

		for j := 0; j < len(nums); j++ {
			if i != j && x+1 == nums[j] {
				j = 0
				x += 1
				count++
			}
		}
		longest = max(longest, count)
	}

	return longest
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	println("longestConsecutiveBruteForce", longestConsecutiveBruteForce(nums))
}
