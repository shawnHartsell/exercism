package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

//Valid checks if the input is a luhn number
func Valid1(input string) bool {
	if len(strings.Trim(input, " ")) <= 1 {
		return false
	}

	matched, _ := regexp.MatchString("^[0-9 ]+$", input)
	if !matched {
		return false
	}

	sum := 0
	double := false
	for i := len(input) - 1; i >= 0; i-- {
		r := input[i]

		if r == ' ' {
			continue
		}

		n, _ := strconv.Atoi(string(r))

		if double {
			n = n * 2
			if n > 9 {
				n = n - 9
			}
		}

		double = !double
		sum += n
	}

	if sum%10 != 0 {
		return false
	}

	return true
}
