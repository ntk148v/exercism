package lsproduct

import (
	"fmt"
	"unicode"
)

// LargestSeriesProduct calculates the largest product for a contigous substring
// of the given digits of length n.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, fmt.Errorf("span is negative: %d", span)
	}
	if len(digits) < span {
		return 0, fmt.Errorf("len(%s) < span: %d < %d", digits, len(digits), span)
	}
	sd := make([]int64, len(digits))
	lsp := int64(0)
	for i, r := range digits {
		if !unicode.IsDigit(r) {
			return 0, fmt.Errorf("digit contains non-digit %q", r)
		}
		sd[i] = int64(r - '0')
	}
	for i := 0; i < len(sd)-span+1; i++ {
		sp := int64(1)
		for _, d := range sd[i : i+span] {
			sp *= d
		}
		if sp > lsp {
			lsp = sp
		}
	}
	return lsp, nil
}
