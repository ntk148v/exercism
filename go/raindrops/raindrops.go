/*Package raindrops converts a number to string,
the contents of which depend on the number's factors.*/
package raindrops

import (
	"strconv"
)

// Convert a number to string
func Convert(number int) (sound string) {
	if number%3 == 0 {
		sound += "Pling"
	}

	if number%5 == 0 {
		sound += "Plang"
	}

	if number%7 == 0 {
		sound += "Plong"
	}

	if len(sound) == 0 {
		sound = strconv.Itoa(number)
	}

	return sound
}
