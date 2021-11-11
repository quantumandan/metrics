package main

import (
	"fmt"

	"github.com/quantumandan/metrics/levenshtein"

	"github.com/quantumandan/metrics/hamming"
)

// Main function for dan's edit distance adventure in golang
func main() {
    fmt.Println("Compare and contrast Levenshtein vs Hamming metrics")

    fmt.Println(levenshtein.EditDistance("intention", "execution"))  // 5
    fmt.Println(hamming.EditDistance("intention", "execution"))      // 5

    fmt.Println(levenshtein.EditDistance("horse", "ros"))            // 3
    fmt.Println(hamming.EditDistance("horse", "ros"))                // 0
    
    fmt.Println(levenshtein.EditDistance("karolin", "kathrin"))      // 3
    fmt.Println(hamming.EditDistance("karolin", "kathrin"))          // 3

    fmt.Println(levenshtein.EditDistance("kitten", "sitting"))       // 3
    fmt.Println(hamming.EditDistance("kitten", "sitting"))           // 0

    // would be really interesting to examing the algebraic and topological
    // connections of the Levenshtein edit distance matrix D and compare it
    // to the equivalent in Hamming for n-length strings.
}
