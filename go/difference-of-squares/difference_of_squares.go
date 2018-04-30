package diffsquares

import "math"

// SquareOfSums returns (1 + ...n)^2
func SquareOfSums(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}

	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares returns (1^2) + ...(n^2)
func SumOfSquares(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2))
	}

	return sum
}

// Difference returns SquareOfSums - SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
