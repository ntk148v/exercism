package anagram

import (
	"sort"
	"strings"
)

// Detect returns a list of correct anagram sublist.
func Detect(subject string, candidates []string) []string {
	var anagrams []string
	for _, c := range candidates {
		if (len(c) != len(subject)) || (strings.ToLower(c) == strings.ToLower(subject)) {
			continue
		}
		if sortString(c) == sortString(subject) {
			anagrams = append(anagrams, c)
		}
	}
	return anagrams
}

// sortString sorts characters in a given string.
func sortString(w string) string {
	s := strings.Split(strings.ToLower(w), "")
	sort.Strings(s)
	return strings.Join(s, "")
}
