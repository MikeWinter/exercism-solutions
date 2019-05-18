// Package scrabble provides functions for calculating word scores
package scrabble

import "strings"

// Score returns the score of word based on standard Scrabble values.
// No multipliers are included.
func Score(word string) int {
	total := 0
	for _, letter := range strings.ToUpper(word) {
		if value, ok := values[letter]; ok {
			total += value
		} else {
			total++
		}
	}
	return total
}

var values = map[rune]int{
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

const (
	defaultValue = 1
)
