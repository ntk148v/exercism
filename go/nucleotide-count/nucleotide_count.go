package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a list of nucleotides.
type DNA string

const validNucleotides = "ACGT"

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var total int
	h := make(Histogram)
	for i := range validNucleotides {
		nucleotide := validNucleotides[i]
		h[nucleotide] = strings.Count(string(d), string(nucleotide))
		total += h[nucleotide]
	}
	if total != len(d) {
		return nil, errors.New("dna: contains invalid nucleotide")
	}
	return h, nil
}
