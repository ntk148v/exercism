/*  Hamming difference between two DNA strands */
package hamming

import "errors"

// Calculate the Hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	diff := 0
	if len(a) != len(b) {
		return diff, errors.New("Two inputs don't have the same length")
	}
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
