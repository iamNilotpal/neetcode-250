package main

import (
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	minLength := min(len(strs[0]), len(strs[len(strs)-1]))

	var builder strings.Builder
	for i := range minLength {
		if strs[0][i] == strs[len(strs)-1][i] {
			builder.Write([]byte{strs[0][i]})
		} else {
			break
		}
	}

	return builder.String()
}

func main() {
	println("longestCommonPrefix", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	println("longestCommonPrefix", longestCommonPrefix([]string{"dog", "race", "car"}))
}
