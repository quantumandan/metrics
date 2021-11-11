package main

import (
	"fmt"

	"github.com/quantumandan/metrics/levenshtein"
)

func main() {
    fmt.Println(levenshtein.EditDistance("intention", "execution"))
}
