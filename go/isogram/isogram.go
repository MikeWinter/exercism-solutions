package isogram

import "unicode"

func IsIsogram(phrase string) bool {
	seen := make(map[rune]bool)

	for _, char := range phrase {
		if char == ' ' || char == '-' {
			continue
		}

		letter := unicode.ToLower(char)
		if seen[letter] {
			return false
		}
		seen[letter] = true
	}
	return true
}
