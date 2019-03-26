package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// Encode composes secrets messages called a square code
func Encode(in string) (out string) {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	in = strings.ToLower(reg.ReplaceAllString(in, ""))
	numCols := int(math.Ceil(math.Sqrt(float64(len(in)))))
	padding := numCols*(numCols-1) - len(in)
	if padding < 0 {
		padding = numCols*numCols - len(in)
	}
	cols := make([]string, numCols)
	for i, r := range in {
		cols[i%numCols] += string(r)
	}
	for i := 0; i < padding; i++ {
		cols[numCols-i-1] += " "
	}
	return strings.Join(cols, " ")
}
