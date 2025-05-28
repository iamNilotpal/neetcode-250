package main

import (
	"slices"
)

func groupAnagramsSol1(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}

	frequency := make(map[string][]string)
	for _, s := range strs {
		sortedStr := sortString(s)
		frequency[sortedStr] = append(frequency[sortedStr], s)
	}

	ans := make([][]string, 0, len(frequency))
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
