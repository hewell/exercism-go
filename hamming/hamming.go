// Package hamming calculates the Hamming difference between two DNA strands.
package hamming

import "errors"

// Distance calculates the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("the two DNA strands are different in length.")
	}
	dist := 0
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
