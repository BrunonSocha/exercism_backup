package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	res := Frequency{}
	re := regexp.MustCompile(`[a-zA-Z0-9]+(?:'[a-zA-Z0-9]+)*`)
	wordArr := re.FindAllString(phrase, -1)
	for _, v := range wordArr {
		if v != "" {
		res[strings.ToLower(v)]++
		}
	}
	return res
}
