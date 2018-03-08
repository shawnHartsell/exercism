package raindrops

import (
	"bytes"
	"strconv"
)

// Convert takes a number and returns a string based on the
// numbers factors

func Convert(n int) string {
	var buffer bytes.Buffer

	if n%3 == 0 {
		buffer.WriteString("Pling")
	}

	if n%5 == 0 {
		buffer.WriteString("Plang")
	}

	if n%7 == 0 {
		buffer.WriteString("Plong")
	}

	if buffer.Len() != 0 {
		return buffer.String()
	}

	return strconv.Itoa(n)
}
