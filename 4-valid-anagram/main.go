package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sFrequency := make(map[rune]int)
	for _, char := range s {
		sFrequency[char] = sFrequency[char] + 1
	}

	for _, char := range t {
		if val, ok := sFrequency[char]; ok && val > 0 {
			sFrequency[char] = val - 1
		}
	}

	for _, sVal := range sFrequency {
		if sVal != 0 {
			return false
		}
	}

	return true
}

func main() {
	println("isAnagram:", isAnagram("", ""))
}
