package romannumerals

import (
	"fmt"
	"strings"
)

func ToRomanNumeral(num int) (string, error) {
	if num <= 0 || num > 3000 {
		return "", fmt.Errorf("The number %d is unable to convert to roman numeral", num)
	}
	var roman strings.Builder
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	for _, conv := range conversions {
		for num >= conv.value {
			roman.WriteString(conv.digit)
			num -= conv.value
		}
	}
	return roman.String(), nil
}
