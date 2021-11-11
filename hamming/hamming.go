package hamming

import "github.com/quantumandan/metrics/utils"

// EditDistance calculate the edit distance between two given strings using only substitutions
// Return edit distance integer
func EditDistance(a, b string) int {
	// Convert strings to rune array to handle no-ASCII characters
	runeA := []rune(a)
	runeB := []rune(b)

	// Compare rune lengths
	if len(runeA) != len(runeB) {
		return 0
	}

	// Check equality of types
	if utils.Equal(runeA, runeB) {
		return 0
	}

	var counter int
	for i, e := range runeA {
		if e != runeB[i] {
			counter++
		}
	}

	return counter
}