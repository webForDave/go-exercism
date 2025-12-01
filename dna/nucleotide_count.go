package dna

import (
	"errors"
	"strings"
)

type Histogram map[rune]int

type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
//
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'T': 0, 'G': 0, 'C': 0,}

	for _, v := range d {
		
		if v != 'G' && v != 'A' && v != 'T' && v != 'C' {
			return nil, errors.New("invalid nucleotide")
		}

		h[v] = strings.Count(string(d), string(v))

	}
	return h, nil
}