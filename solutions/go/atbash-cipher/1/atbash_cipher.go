package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	plain := "abcdefghijklmnopqrstuvwxyz"
	cipher := "zyxwvutsrqponmlkjihgfedcba"
	s = strings.ToLower(s)
	res := ""
	groupCount := 0
	for _, v := range s {
		if groupCount == 5 {
			res += " "
			groupCount = 0
		}
		if unicode.IsDigit(v) {
			res += string(v)
			groupCount++
			continue
		}
		for i, l := range cipher {
			if v == l {
				res += string(plain[i])
				groupCount++
			} 	
		}
	}
	return strings.TrimSpace(res)
}
