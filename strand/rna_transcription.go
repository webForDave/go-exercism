package strand

import "strings"

func ToRNA(dna string) string {

	var complement strings.Builder

	for _, v := range dna {
		switch {
		case v == 'G':
			complement.WriteString("C")
		case v == 'C':
			complement.WriteString("G")
		case v == 'A':
			complement.WriteString("U")
		case v == 'T':
			complement.WriteString("A")
		}
	}
	return complement.String()
}