package main

import (
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	// If the input array of strings is empty, there's no common prefix.
	if len(strs) == 0 {
		return ""
	}
	// If there's only one string, that string itself is the longest common prefix.
	if len(strs) == 1 {
		return strs[0]
	}

	// Intuition 2: The LCP must be a common prefix of the "extremes".
	// When strings are sorted lexicographically:
	// - The first string (strs[0]) will be the "smallest" string.
	// - The last string (strs[len(strs)-1]) will be the "largest" string.
	// Any common prefix among *all* strings must also be a common prefix
	// of the first and last strings after sorting. If the first and last
	// strings share a prefix, all strings in between (lexicographically)
	// *must* also share that same prefix.
	sort.Strings(strs)

	// Intuition 3: Determine the maximum possible length of the LCP.
	// The longest common prefix cannot be longer than the shortest of the two extreme strings.
	minLength := len(strs[0])

	// Intuition 4: Build the common prefix character by character.
	// Use a strings.Builder for efficient string concatenation, as opposed to `+=`.
	var builder strings.Builder
	// Iterate character by character up to the determined 'minLength'.
	for i := range minLength {
		// Compare the character at the current index `i` from the first and last strings.
		if strs[0][i] == strs[len(strs)-1][i] {
			// If they match, this character is part of the common prefix. Append it.
			builder.WriteByte(strs[0][i]) // Use WriteByte for single bytes (characters).
		} else {
			// If characters at the current index `i` do not match,
			// it means we've found the end of the common prefix.
			// Break the loop, as no further characters can be part of the LCP.
			break
		}
	}

	return builder.String()
}

func main() {
	println("longestCommonPrefix", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	println("longestCommonPrefix", longestCommonPrefix([]string{"dog", "race", "car"}))
}
