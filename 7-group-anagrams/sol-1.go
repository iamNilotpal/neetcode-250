package main

import (
	"slices"
)

func groupAnagramsSol1(strs []string) [][]string {
	// Intuition 1: Handle Edge Cases
	// If there's 0 or 1 string, there's nothing to group or only one group.
	if len(strs) <= 1 {
		return [][]string{strs}
	}

	// Intuition 2: Use a Hash Map to Group Anagrams
	// The core idea is that anagrams, when their characters are sorted, will result
	// in the exact same sorted string. This sorted string can serve as a unique "key"
	// for all its anagrams.
	// The map will store: key = sorted string (canonical form), value = slice of original strings that are anagrams.
	frequency := make(map[string][]string)

	// Intuition 3: Iterate and Populate the Map
	// For each original string `s` in the input `strs`:
	for _, s := range strs {
		// 1. Sort the characters of the current string `s`.
		// This `sortedStr` is the "canonical form" or "signature" of the anagram group.
		sortedStr := sortString(s)

		// 2. Append the original string `s` to the list associated with its `sortedStr` key.
		// If `sortedStr` is a new key, a new empty slice will be implicitly created.
		// Otherwise, `s` is appended to the existing slice for that key.
		frequency[sortedStr] = append(frequency[sortedStr], s)
	}

	// Intuition 4: Collect the Grouped Anagrams
	// The map now contains all the anagrams grouped by their sorted forms.
	// We need to convert this map structure into the required `[][]string` format.
	// Initialize a result slice. Pre-allocating capacity based on map size can be a minor optimization.
	ans := make([][]string, 0, len(frequency))
	// Iterate through the values of the `frequency` map.
	// Each `val` here is a `[]string` containing a group of anagrams.
	for _, val := range frequency {
		ans = append(ans, val)
	}

	return ans
}

func sortString(str string) string {
	runes := []rune(str)
	slices.Sort(runes)
	return string(runes)
}
