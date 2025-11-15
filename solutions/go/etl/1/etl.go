package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := make(map[string]int)
	for v, p := range in {
		for _, l := range p {
			res[strings.ToLower(l)] = v
		}
	}
	return res
}

