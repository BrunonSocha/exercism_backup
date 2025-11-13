package scrabble

import "strings"

func Score(word string) int {
	res := 0
	for _, l := range strings.ToUpper(word) {
		switch l {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			res++
		case 'D', 'G':
			res += 2
		case 'B', 'C', 'M', 'P':
			res += 3
		case 'F', 'H', 'V', 'W', 'Y':
			res += 4
		case 'K':
			res += 5
		case 'J', 'X':
			res += 8
		case 'Q', 'Z':
			res += 10
		}
	}
	return res
}
