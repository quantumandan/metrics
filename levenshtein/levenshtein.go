package levenshtein

import "github.com/quantumandan/metrics/utils"

// EditDistance calculate the distance between two string
// This algorithm allow insertions, deletions and substitutions to change one string to the second
// Compatible with non-ASCII characters
func EditDistance(a, b string) int {
	// Convert string parameters to rune arrays to be compatible with non-ASCII
	runeA := []rune(a)
	runeB := []rune(b)

	// Get and store length of these strings
	aLen := len(runeA)
	bLen := len(runeB)
	if aLen == 0 {
		return bLen
	} else if bLen == 0 {
		return aLen
	} else if utils.Equal(runeA, runeB) {
		return 0
	}

	dSlice := make([]int, aLen + 1)

	for col := 1; col <= aLen; col++ {
		dSlice[col] = col
	}

	for x := 1; x <= bLen; x++ {
		dSlice[0] = x
		lastkey := x - 1
		for y := 1; y <= aLen; y++ {
			oldkey := dSlice[y]
			var i int
			if runeA[y-1] != runeB[x-1] {
				i = 1
			}
			dSlice[y] = minimum(
							minimum(dSlice[y] + 1, // insert
							dSlice[y - 1] + 1),    // delete
							lastkey + i)           // substitution
			lastkey = oldkey
		}
	}

	return dSlice[aLen]
}

func minimum(x, y int) int {
// Minimum returns the smallest integer among the two in parameters
	if y < x {
        return y
    }
	return x
}
