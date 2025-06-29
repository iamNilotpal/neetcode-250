package main

import "math"

func scoreOfString(s string) int {
	if len(s) <= 1 {
		return 0
	}

	var score int
	for i := range len(s) - 1 {
		score += int(math.Abs(float64(int(s[i]) - int(s[i+1]))))
	}

	return score
}

func main() {
	println("Score of a String", scoreOfString("hello"))
}
