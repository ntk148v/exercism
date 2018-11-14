package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid validates a variety of identification.
func Valid(input string) bool {
	// Remove all whitespace
	input = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, input)
	if _, err := strconv.Atoi(input); err != nil || len(input) <= 1 {
		return false
	}
	var sum int
	for i, v := range input {
		// Use j as reverse index
		j := len(input) - i
		v -= '0'
		sum += int(v)
		if j%2 == 0 {
			sum += int(v)
			if int(v)*2 > 9 {
				sum -= 9
			}
		}
	}
	return sum%10 == 0
}
