package levenshtein

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
	} else if Equal(runeA, runeB) {
		return 0
	}

	dSlice := make([]int, aLen + 1)

	for y := 1; y <= aLen; y++ {
		dSlice[y] = y
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
			dSlice[y] = Minimum(
							Minimum(dSlice[y] + 1, // insert
							dSlice[y - 1] + 1),    // delete
							lastkey + i)           // substitution
			lastkey = oldkey
		}
	}

	return dSlice[aLen]
}

func Minimum(x, y int) int {
// Minimum returns the smallest integer among the two in parameters
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