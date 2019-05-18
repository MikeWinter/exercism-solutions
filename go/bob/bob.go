package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == `` {
		return "Fine. Be that way!"
	}

	question := strings.HasSuffix(remark, `?`)
	calm := false
	statement := false
	for _, char := range remark {
		if !calm && unicode.IsLower(rune(char)) {
			calm = true
		}
		if !statement && unicode.IsLetter(rune(char)) {
			statement = true
		}
	}
	shouting := statement && !calm

	switch {
	case question && shouting:
		return "Calm down, I know what I'm doing!"
	case question:
		return "Sure."
	case shouting:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
