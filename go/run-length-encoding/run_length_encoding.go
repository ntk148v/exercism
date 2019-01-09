package encode

import (
	"fmt"
	"strconv"
	"strings"
)

// RunLengthEncode encodes the given string
func RunLengthEncode(in string) string {
	count := 1
	var output string
	for i, r := range in {
		if i != 0 {
			if rune(in[i-1]) == r {
				count++
			} else {
				if count > 1 {
					output += fmt.Sprintf("%d%c", count, in[i-1])
				} else {
					output += fmt.Sprintf("%c", in[i-1])
				}
				count = 1
			}
		}
	}
	if len(in) != 0 {
		if count > 1 {
			output += fmt.Sprintf("%d%c", count, in[len(in)-1])
		} else {
			output += fmt.Sprintf("%c", in[len(in)-1])
		}
	}
	return output
}

// RunLengthDecode decodes the given string
func RunLengthDecode(in string) string {
	count := 1
	var stringCount, output string
	for _, r := range in {
		if _, err := strconv.Atoi(string(r)); err == nil {
			stringCount += string(r)
		} else {
			if stringCount != "" {
				count, _ = strconv.Atoi(stringCount)
			}
			output += strings.Repeat(string(r), count)
			count = 1
			stringCount = ""
		}
	}
	return output
}
