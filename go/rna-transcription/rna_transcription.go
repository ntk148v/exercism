package strand

import "strings"

// ToRNA transcribes a given DNA strand to RNA strand.
func ToRNA(dna string) string {
	var builder strings.Builder
	dna = strings.ToUpper(dna)
	for _, value := range dna {
		switch string(value) {
		case "G":
			builder.WriteString("C")
		case "C":
			builder.WriteString("G")
		case "T":
			builder.WriteString("A")
		case "A":
			builder.WriteString("U")
		}
	}
	return builder.String()
}
