package main

import (
	"math"
)

func hasAllCodes(s string, k int) bool {
	table := make(map[string]bool)
	count := int(math.Pow(2, float64(k)))
	for i := 0; i < len(s)-k+1; i++ {
		table[s[i:i+k]] = true
		if len(table) == count {
			return true
		}
	}
	return len(table) == count
}

func main() {
}
