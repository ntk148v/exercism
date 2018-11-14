package protein

import "errors"

var (
	STOP           = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")
)

// FromCodon returns the protein for the given codon.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", STOP
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA returns the protein sequence for an RNA strand.
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	bases := []rune(rna)
	for i := 0; i < len(bases); i += 3 {
		p, err := FromCodon(string(bases[i : i+3]))
		switch err {
		case STOP:
			return proteins, nil
		case ErrInvalidBase:
			return proteins, err
		default:
			proteins = append(proteins, p)
		}
	}
	return proteins, nil
}
