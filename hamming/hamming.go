// Package hamming provides a function to calculate the Hamming distance of 2 DNA strands.
package hamming

import "errors"

// Distance expects two strings as a representation of two strands. This is only possible
// if these are equal in length. If they are the Hamming distance is returned as an integer
// , if not the function returns an error.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("sequences need to have the same length")
	}
	var dist int
	// Iterating over the range of strand A and comparing the rune to the same position in strand B
	for i, nucleotideA := range a {
		if nucleotideA != rune(b[i]) {
			dist++
		}
	}

	return dist, nil
}
