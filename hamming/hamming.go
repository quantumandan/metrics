package hamming

import "github.com/quantumandan/metrics/utils"

// HammingDistance calculate the edit distance between two given strings using only substitutions
// Return edit distance integer
func EditDistance(a, b string) int {
	// Convert strings to rune array to handle no-ASCII characters
	runeA := []rune(a)
	runeB := []rune(b)

	if len(runeA) != len(runeB) {
		return 0
	} else if utils.Equal(runeA, runeB) {
		return 0
	}

	var counter int
	for i := 0; i < len(runeA); i++ {
		if runeA[i] != runeB[i] {
			counter++
		}
	}

	return counter
}
