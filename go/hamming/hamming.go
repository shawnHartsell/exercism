package hamming

import (
	"errors"
)

// Distance caclulates the Hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strings must be the same length")
	}

	if a == "" && b == "" {
		return 0, nil
	}

	d := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}

	return d, nil
}
