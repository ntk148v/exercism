package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(input string) bool {
	input = strings.ToLower(input)
	for index, character := range input {
		if unicode.IsLetter(character) {
			if lastIndex := strings.LastIndexAny(input, string(character)); lastIndex > index {
				return false
			}
		}
	}
	return true
}
