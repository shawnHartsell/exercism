//Package bob represents a lackadaisacal teenage
package bob

import (
	"strings"
	"unicode"
)

// Hey takes a remark and returns bob's reply
func Hey(remark string) string {
	var answer string

	remark = strings.Trim(remark, "\t\n\r ")
	isQuestion := strings.HasSuffix(remark, "?")
	isYelling := isYelling(remark)

	switch {
	case remark == "":
		answer = "Fine. Be that way!"
	case isYelling && isQuestion:
		answer = "Calm down, I know what I'm doing!"
		break
	case isQuestion:
		answer = "Sure."
		break
	case isYelling:
		answer = "Whoa, chill out!"
		break
	default:
		answer = "Whatever."
	}

	return answer
}

func isYelling(remark string) bool {
	var hasChar bool
	isYelling := true

	for _, c := range remark {
		if unicode.IsLetter(c) {
			hasChar = true

			if unicode.IsLower(c) {
				isYelling = false
				break
			}
		}
	}

	if !hasChar {
		return false
	}

	return isYelling
}
