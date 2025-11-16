package anagram

import (
	"strings"
	"slices"
)

func Detect(subject string, candidates []string) []string {
	var res []string
	sArr := []rune(strings.ToLower(subject))
	slices.Sort(sArr)
	for _, c := range candidates {
		anagram := true
		if strings.ToLower(c) == strings.ToLower(subject) || len(c) != len(subject) {
			continue
		}
		cArr := []rune(strings.ToLower(c))
		slices.Sort(cArr)
		for i := 0; i < len(sArr); i++ {
			if cArr[i] != sArr[i] {
				anagram = false
			}
		}
		if anagram {
			res = append(res, c)
		}
	}	
	return res
}
