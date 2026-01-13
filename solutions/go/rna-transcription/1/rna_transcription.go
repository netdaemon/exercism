package strand

var dnaToRna = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	rna := ""

	for _, n := range dna {
		rna += string(dnaToRna[n])
	}

	return rna
}
