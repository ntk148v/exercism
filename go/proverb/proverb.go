package proverb

import (
	"fmt"
	"strings"
)

// Proverb generate the relevant proverbs with a given list of inputs.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	var proverbs []string
	if len(rhyme) > 1 {
		for i := 0; i < len(rhyme)-1; i++ {
			proverbs = append(proverbs, fmt.Sprintf("For want of a %s the %s was lost.", strings.ToLower(rhyme[i]), strings.ToLower(rhyme[i+1])))
		}
	}
	proverbs = append(proverbs, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return proverbs
}
