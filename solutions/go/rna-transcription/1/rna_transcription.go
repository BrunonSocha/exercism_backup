package strand

func ToRNA(dna string) string {
	rnaconv := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	var res []rune
	for _, v := range dna {
		res = append(res, rnaconv[v])
	}
	return string(res)
}
