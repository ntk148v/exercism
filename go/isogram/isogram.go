package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(input string) bool {
	input = strings.ToLower(input)
	var count int
	for _, character := range input {
		if unicode.IsLetter(character) {
			count = strings.Count(input, string(character))
			if count > 1 {
				return false
			}
		}
	}
	return true
}
