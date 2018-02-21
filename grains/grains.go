// Package grains calculate the number of grains on the chessboard
package grains

import "errors"
import "math"

// Square calculates the number of grains in a given square
func Square(in int) (uint64, error) {
	if in <= 0 {
		return 0, errors.New("Square number must be greater than 1.")
	}
	if in > 64 {
		return 0, errors.New("Square number must be no greater than 64.")
	}
	return uint64(math.Exp2(float64(in) - 1)), nil
}

// Total returns the total number of grains on the chessboard
func Total() uint64 {
	return 18446744073709551615
}
