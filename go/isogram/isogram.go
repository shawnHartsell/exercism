package isogram

import "unicode"

// IsIsogram determines if a string has repeating characters
func IsIsogram(s string) bool {
	if s == "" {
		return true
	}

	m := make(map[rune]bool)

	for _, c := range s {
		if c == '-' || c == ' ' {
			continue
		}

		if m[c] {
			return false
		}

		m[unicode.ToLower(c)] = true
	}

	return true
}
