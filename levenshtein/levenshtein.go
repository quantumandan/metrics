package levenshtein

// Return the smallest integer among the two in parameters
func minimum(a, b int) int {
    if b < a {
        return b
    }
	return a
}

func EditDistance(a, b string) int {
	// Convert strings to runes (less brittle)
	runeWord1 := []rune(a)
	runeWord2 := []rune(b)

	// Dimensions of the edit distance matrix
	m := len(runeWord1) + 1  // number columns
	n := len(runeWord2) + 1  // number rows

	// Create edit distance matrix D
	D := [m][n]int{}

	// Initialize the columns
	for i := 1; i <= m; i++ {
		D[i][0] = i
	}
	// Initialize the rows
	for j := 1; j <= n; j++ {
		D[0][j] = j
	}

	// Calculate the cost of a substitution
	// for j := 1; j <= n; j++ {
	// 	for i := 1; i <= m; i++ {
	// 		subcost := 0
	// 		if (runeWord1[i] != runeWord2[j]) {
	// 			subcost = 1
	// 		}

	// 		// editCosts := []int{D[i - 1][j] + 1, D[i][j - 1] + 1, D[i - 1][j - 1] + subcost}
	// 		// min := editCosts[0]
	// 		// for _, v := range editCosts {
	// 		// 	if (v < min) {
	// 		// 		min = v
	// 		// 	}
	// 		// }

	// 		D[i][j] = minimum(
	// 					minimum(D[i - 1][j] + 1, D[i][j - 1] + 1),
	// 							D[i - 1][j - 1] + subcost)

	// 		// D[i][j] = minimum(
	// 		// 			minimum(
	// 		// 				D[i - 1][j] + 1,  		// deletion
	// 		// 				D[i][j - 1] + 1),  		// insertion
	// 		// 				D[i - 1][j - 1] + cost)  // substitution
					
	// }
	return D[m][n]
}