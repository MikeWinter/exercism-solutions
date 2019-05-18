package acronym

import (
	"regexp"
	"strings"
)

var wordPattern *regexp.Regexp

func init() {
	wordPattern = regexp.MustCompile(`([A-Za-z])[\w']*`)
}

func Abbreviate(s string) string {
	words := wordPattern.FindAllString(s, -1)
	acronym := ``
	for _, word := range words {
		acronym += strings.ToUpper(string(word[0]))
	}
	return acronym
}
