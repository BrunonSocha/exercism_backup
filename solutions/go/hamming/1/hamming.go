package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	aMap := make(map[int]rune)
	bMap := make(map[int]rune)
	res := 0
	if len(a) != len(b) {
		return 0, errors.New("The strands must be the same length.") 
	}
	for i, v := range a {
		aMap[i] = v
	}
	for i, v := range b {
		bMap[i] = v
	}
	for i, _ := range a {
		if aMap[i] != bMap[i] {
			res += 1
		}
	}
	return res, nil
}
