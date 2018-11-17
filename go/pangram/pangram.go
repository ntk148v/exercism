package pangram

import "unicode"

// IsPangram returns whether a sentence is a pangram or not.
func IsPangram(sen string) bool {
	seen := make(map[rune]struct{}, len(sen))
	// Remove duplicated letters
	for _, r := range sen {
		if _, ok := seen[r]; ok || !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		seen[r] = struct{}{}
	}
	return len(seen) == 26
}
