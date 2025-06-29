package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	// Step 1: Calculate the frequency of each number.
	frequency := make(map[int]int)
	for _, v := range nums {
		frequency[v] += 1
	}

	// Step 2: Create "buckets" based on frequency.
	// This is the core idea of Bucket Sort or a variation of it.
	// The index of the slice will represent the frequency, and the value at that index
	// will be a slice of numbers that have that specific frequency.
	buckets := make([][]int, len(nums)+1)
	for i := range buckets {
		buckets[i] = make([]int, 0, 1)
	}

	// Step 3: Populate the buckets.
	// Iterate through the `frequency` map.
	// For each number (key) and its count (val):
	// Place the number into the bucket corresponding to its count.
	for key, val := range frequency {
		buckets[val] = append(buckets[val], key) // `val` is the frequency, `key` is the number.
	}

	// Step 4: Collect the top K frequent elements from the buckets.
	// We want the most frequent elements, so we iterate the buckets from the highest frequency
	// down to the lowest (from `len(buckets) - 1` down to 0).
	var result []int
	for i := len(buckets) - 1; i >= 0; i-- {
		bucket := buckets[i] // Get the slice of numbers at the current frequency `i`.

		// If the current bucket is empty, there are no numbers with this frequency, so skip it.
		if len(bucket) == 0 {
			continue
		}

		// Append all numbers from the current bucket to our `result` slice.
		// Since we're iterating from highest frequency down, numbers added here are more frequent.
		for _, v := range bucket {
			result = append(result, v)
		}

		// Optimization: If we have already collected `k` or more elements,
		// we have found our top K. We can stop early.
		if len(result) >= k {
			break
		}
	}

	// Step 5: Return the first K elements.
	// In case we collected more than `k` elements in the last bucket (e.g., multiple numbers
	// share the same K-th highest frequency), we slice the `result` to return exactly `k` elements.
	return result[:k]
}

func main() {
	array := []int{1, 1, 1, 2, 2, 3}
	fmt.Printf("topKFrequent: %+v \n", topKFrequent(array, 2))
}
