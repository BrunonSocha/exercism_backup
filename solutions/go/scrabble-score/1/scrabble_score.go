package scrabble

import "strings"

func Score(word string) int {
	scoreTable := map[int]([]rune){1: {'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'}, 2: {'D', 'G'}, 3: {'B', 'C', 'M', 'P'}, 4: {'F', 'H', 'V', 'W', 'Y'}, 5: {'K'}, 8: {'J', 'X'}, 10: {'Q', 'Z'}}
	res := 0
	wordrunes := []rune(strings.ToUpper(word))
	for ind := 0; ind < len(wordrunes); ind++ {
		for i, v := range scoreTable {
			for _, cl := range v {
				if wordrunes[ind] == cl {
					res += i
				}
			}
		} 
	}
	return res
}
