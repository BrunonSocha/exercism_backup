package protein

import "errors"

var (
	ErrStop = errors.New("STOP")
	ErrInvalidBase = errors.New("Incorrect sequence")
)

func FromRNA(rna string) ([]string, error) {
	var res []string
	buffer := ""
	for _, v := range rna {
		buffer += string(v)
		if len(buffer) == 3 {
			protein, err := FromCodon(buffer)
			if err == ErrStop {
				break
			} else if err == ErrInvalidBase {
				return res, err
			}
			res = append(res, protein)
			buffer = ""
		}
	}
	return res, nil
}

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
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
