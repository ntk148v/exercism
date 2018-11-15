package reverse

import "strings"

// Reverse returns a reverse of given string.
func String(in string) string {
	lenIn := len(in)
	out := make([]rune, lenIn)
	for i, r := range in {
		out[lenIn-1-i] = r
	}
	return strings.Replace(string(out), "\x00", "", -1)
}
