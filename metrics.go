package main

import (
	"fmt"

	"github.com/quantumandan/metrics/levenshtein"

	"github.com/quantumandan/metrics/hamming"
)

func main() {
    fmt.Println(levenshtein.EditDistance("intention", "execution"))
    fmt.Println(hamming.EditDistance("intention", "execution"))
    fmt.Println(levenshtein.EditDistance("horse", "ros"))
    fmt.Println(hamming.EditDistance("horse", "ros"))
}
