package levenshtein

// EditDistance calculate the distance between two string
// This algorithm allow insertions, deletions and substitutions to change one string to the second
// Compatible with non-ASCII characters
func EditDistance(a, b string) int {
	// Convert string parameters to rune arrays to be compatible with non-ASCII
	runeA := []rune(a)
	runeB := []rune(b)

	// Get and store length of these strings
	ALen := len(runeA)
	BLen := len(runeB)
	if ALen == 0 {
		return BLen
	} else if BLen == 0 {
		return ALen
	} else if Equal(runeA, runeB) {
		return 0
	}

	DSlice := make([]int, ALen + 1)

	for y := 1; y <= ALen; y++ {
		DSlice[y] = y
	}
	for x := 1; x <= BLen; x++ {
		DSlice[0] = x
		lastkey := x - 1
		for y := 1; y <= ALen; y++ {
			oldkey := DSlice[y]
			var i int
			if runeA[y-1] != runeB[x-1] {
				i = 1
			}
			DSlice[y] = Minimum(
							Minimum(DSlice[y] + 1, // insert
							DSlice[y - 1] + 1),    // delete
							lastkey + i)           // substitution
			lastkey = oldkey
		}
	}

	return DSlice[ALen]
}

func Minimum(x, y int) int {
// Return the smallest integer among the two in parameters
	if y < x {
        return y
    }
	return x
}

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