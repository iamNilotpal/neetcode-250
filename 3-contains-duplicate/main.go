package main

func containsDuplicate(nums []int) bool {
	// Intuition: Use a hash map (or frequency map) to keep track of elements seen so far.
	// If we encounter an element that's already in our map, it means we've found a duplicate.

	// 1. Create a hash map:
	// 'frequency' will store each unique number encountered as a key,
	// and its count (though we only truly care if it's > 0) as the value.
	frequency := make(map[int]int)

	// 2. Iterate through the input array:
	// For each 'item' (number) in 'nums':
	for _, item := range nums {
		// 3. Check for existence in the map:
		// 'data' will get the current count of 'item' from the map.
		// If 'item' is not in the map, 'data' will be the zero value for int, which is 0.
		data := frequency[item]

		// 4. If 'data' is greater than 0, it means 'item' has been seen before.
		// Therefore, we've found a duplicate, and we can immediately return true.
		if data > 0 {
			return true
		}

		// 5. If 'item' was not a duplicate (i.e., 'data' was 0),
		// increment its count in the map. This marks it as "seen".
		frequency[item] = data + 1
	}

	// 6. If the loop completes without returning true, it means no duplicates were found.
	// So, return false.
	return false
}

func main() {
	println("Contains Duplicate", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
