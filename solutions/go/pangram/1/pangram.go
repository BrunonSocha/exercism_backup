package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	sum := 0
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, v := range(input) {
		for _, a := range(alphabet) {
			if v == a {
				sum++
				alphabet = strings.ReplaceAll(alphabet, string(a), "")
			}  
		}
	}
	return sum == 26
}
