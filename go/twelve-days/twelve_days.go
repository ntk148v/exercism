package twelve

import (
	"fmt"
)

var dayGifts = [][]string{
	[]string{"first", "a Partridge in a Pear Tree."},
	[]string{"second", "two Turtle Doves"},
	[]string{"third", "three French Hens"},
	[]string{"fourth", "four Calling Birds"},
	[]string{"fifth", "five Gold Rings"},
	[]string{"sixth", "six Geese-a-Laying"},
	[]string{"seventh", "seven Swans-a-Swimming"},
	[]string{"eighth", "eight Maids-a-Milking"},
	[]string{"ninth", "nine Ladies Dancing"},
	[]string{"tenth", "ten Lords-a-Leaping"},
	[]string{"eleventh", "eleven Pipers Piping"},
	[]string{"twelfth", "twelve Drummers Drumming"},
}

const form = "On the %s day of Christmas my true love gave to me%s"

func Verse(i int) string {
	var gifts string
	for j := i - 1; j >= 0; j-- {
		if i != 1 && j == 0 {
			gifts = fmt.Sprintf("%s, and %s", gifts, dayGifts[j][1])
		} else if i != j+1 {
			gifts = fmt.Sprintf("%s, %s", gifts, dayGifts[j][1])
		} else {
			gifts = fmt.Sprintf("%s: %s", gifts, dayGifts[j][1])
		}
	}
	return fmt.Sprintf(form, dayGifts[i-1][0], gifts)
}

func Song() string {
	var song = ""
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song
}
