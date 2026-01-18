package protein

import (
	"errors"
)

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("errInvalidBase")

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}

	for i := 0; i <= len(rna)-3; i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)

		if err == ErrStop {
			return proteins, nil
		}

		if err == ErrInvalidBase {
			return proteins, ErrInvalidBase
		}

		proteins = append(proteins, protein)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	protein, found := codons[codon]

	switch {
	case !found:
		return "", ErrInvalidBase
	case protein == "STOP":
		return protein, ErrStop
	default:
		return protein, nil
	}

}
