//Package acronym abbreviates strings
package acronym

import (
	"strings"
)

/*
Abbreviate takes a string and creates an abbreviaion using the first
character of each word (separated by space).
Hypenated words count as distinct words
*/
func Abbreviate(s string) string {
	abbreviation := ""

	for _, token := range strings.Split(s, " ") {
		abbreviation += strings.ToUpper(string(token[0]))

		subTokens := strings.Split(token, "-")
		if len(subTokens) > 1 {
			abbreviation += strings.ToUpper(string(subTokens[1][0]))
		}

	}

	return abbreviation
}
