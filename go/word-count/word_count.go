package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a map of the frequency of occurrence to the unique word.
type Frequency map[string]int

// WordCount returns the map of frequency of words.
func WordCount(phrase string) Frequency {
	freq := Frequency{}
	words := strings.Fields(normalize(phrase))
	for _, w := range words {
		w = strings.Trim(w, "'")
		freq[w]++
	}
	return freq
}

func normalize(phrase string) string {
	r, _ := regexp.Compile(`[^\w|']`)
	phrase = strings.ToLower(phrase)
	return r.ReplaceAllLiteralString(phrase, " ")
}
