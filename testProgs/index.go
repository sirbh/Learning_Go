package testProgs

// CharCounter counts the occurrences of each character in the string
func CharCounter(b string) map[rune]int {
	charCount := make(map[rune]int)

	for _, ch := range b {
		charCount[ch]++
	}

	return charCount
}
