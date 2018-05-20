package grains

import (
	"errors"
	"math"
)

//Square returns the number of grains that would be on the specified chess square
//if the number of grains doubled
func Square(square int) (uint64, error) {
	if square <= 0 || square > 64 {
		return 0, errors.New("square must be between 1 and 64")
	}

	return uint64(math.Pow(float64(2), float64(square-1))), nil
}

//Total returns the total number of grains on a 64 square
//chessboard if the number is doubled on each square
func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		sum += s
	}

	return sum
}
