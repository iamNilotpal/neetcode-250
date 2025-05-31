package main

func isAnagram(s string, t string) bool {
	// Intuition 1: Anagrams must have the same length.
	// If the lengths are different, they cannot be anagrams. This is a quick exit condition.
	if len(s) != len(t) {
		return false
	}

	// Intuition 2: Count character frequencies for the first string.
	// A hash map (or frequency map) is ideal for storing character counts.
	// 'rune' is used for characters to correctly handle Unicode.
	sFrequency := make(map[rune]int)
	for _, char := range s {
		sFrequency[char] = sFrequency[char] + 1
	}

	// Intuition 3: Decrement counts based on the second string.
	// Now, iterate through the characters of string 't'.
	for _, char := range t {
		// Check if the character exists in our frequency map and if its count is greater than 0.
		// `val, ok := sFrequency[char]` checks if `char` exists as a key (`ok` will be true)
		// and retrieves its current count (`val`).
		if val, ok := sFrequency[char]; ok && val > 0 {
			// If the character is found and its count is positive, decrement its count.
			// This simulates "using up" a character from the pool available from 's'.
			sFrequency[char] = val - 1
		} else {
			// Intuition 3a: If a character from 't' is not found in 'sFrequency' (meaning it wasn't in 's' at all)
			// OR if its count is already 0 (meaning 't' has more of this character than 's'),
			// then 't' cannot be an anagram of 's'. Return false immediately.
			// This is a crucial optimization to fail early.
			return false
		}
	}

	// Intuition 4: Final check for remaining counts.
	// After processing all characters in 't', go through the 'sFrequency' map.
	// If 's' and 't' are perfect anagrams, all character counts in `sFrequency`
	// should now be exactly 0 (meaning every character from 's' was matched and used up by 't').
	for _, sVal := range sFrequency {
		if sVal != 0 {
			// If any character count is not zero, it means 's' had characters
			// that 't' didn't have (count > 0) or 't' had too many of a character
			// that wasn't properly accounted for (though the `else` block above
			// handles this more directly for counts that go negative).
			return false
		}
	}

	// If all checks pass, it means 't' is an anagram of 's'.
	return true
}

func main() {
	println("isAnagram:", isAnagram("", ""))
}
