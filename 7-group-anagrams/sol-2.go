package main

import "strings"

func groupAnagramsSol2(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}

	frequency := make(map[string][]string)
	for _, str := range strs {
		strFrequency := getFrequencyString(str)
		frequency[strFrequency] = append(frequency[strFrequency], str)
	}

	ans := make([][]string, 0, len(frequency))
	for _, v := range frequency {
		ans = append(ans, v)
	}
	return ans
}

func getFrequencyString(str string) string {
	charFrequency := make([]int, 26)
	for _, char := range str {
		index := int(char - 'a')
		charFrequency[index]++
	}

	char := 'a'
	var builder strings.Builder

	for _, v := range charFrequency {
		builder.WriteRune(char)
		builder.WriteByte(byte(v))
		char++
	}

	return builder.String()
}
