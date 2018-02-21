// Package diffsquares calculates difference between
// Square of Sums and Sum of Squares of the first N natural numbers
package diffsquares

// SquareOfSums calculates the square of sums of the first N natural numbers
func SquareOfSums(in int) int {
	var sum = Sum(in)
	return sum * sum
}

// SumOfSquares calculates the sum of squares of the first N natural numbers
func SumOfSquares(in int) int {
	var sum = 0
	for i := 1; i <= in; i++ {
		sum += i * i
	}
	return sum
}

// Difference calculates the difference of SquareOfSums & SumOfSquares
func Difference(in int) int {
	return SquareOfSums(in) - SumOfSquares(in)
}

// Sum calculates the sum of the first N natural numbers
func Sum(in int) int {
	return (1 + in) * in / 2
}
