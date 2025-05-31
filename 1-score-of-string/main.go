package main

import "math"

func scoreOfString(s string) int {
	// If the string has 0 or 1 character, there are no adjacent pairs.
	// Therefore, no differences can be calculated, and the score must be 0.
	if len(s) <= 1 {
		return 0
	}

	var score int
	// We need to compare each character with its immediate neighbor.
	// The loop runs from the first character (index 0) up to the
	// second-to-last character (index len(s) - 2).
	// Why? Because when `i` is `len(s) - 2`, `s[i+1]` will correctly
	// point to the *last* character (`s[len(s) - 1]`), covering all pairs.
	for i := range len(s) - 1 {
		// 1. Get the ASCII values: `s[i]` and `s[i+1]` give us the byte (ASCII) values.
		// 2. Convert to int: `int(s[i])` and `int(s[i+1])` cast these bytes to integers.
		// 3. Find the difference: `int(s[i]) - int(s[i+1])` gives us how much they differ.
		// 4. Take the absolute value: `math.Abs(...)` ensures we always get a positive
		//    "distance" between their ASCII values, regardless of which one is larger.
		// 5. Convert back to int: `int(...)` casts the float64 result of `math.Abs` back to an integer.
		score += int(math.Abs(float64(int(s[i]) - int(s[i+1]))))
	}

	return score
}

func main() {
	println("Score of a String", scoreOfString("hello"))
}
