package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	res := 0
	if len(a) != len(b) {
		return 0, errors.New("The strands must be the same length.") 
	}
	for i, _ := range a {
		if a[i] != b[i] {
			res++
		}
	}
	return res, nil
}
