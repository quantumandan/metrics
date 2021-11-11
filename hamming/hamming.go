package hamming

// Equal compare two rune arrays and return if they are equals or not
func Equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// HammingDistance calculate the edit distance between two given strings using only substitutions
// Return edit distance integer
func EditDistance(a, b string) int {
	// Convert strings to rune array to handle no-ASCII characters
	runeA := []rune(a)
	runeB := []rune(b)

	if len(runeA) != len(runeB) {
		return 0
	} else if Equal(runeA, runeB) {
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
