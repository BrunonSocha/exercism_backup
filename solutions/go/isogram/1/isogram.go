package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(strings.ReplaceAll((strings.ReplaceAll(word, " ", "")), "-", ""))
	for i := 0; i < len(word); i++ {
		for j := i+1; j < len(word); j++{
			if word[i] == word[j] {
				return false
			}
		}
	}
	return true
}
