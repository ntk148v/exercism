package isogram

import "unicode"

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(input string) bool {
	seen := map[rune]bool{}
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
