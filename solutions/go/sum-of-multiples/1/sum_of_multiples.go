package summultiples

import "slices"

func SumMultiples(limit int, divisors ...int) int {
	var listMult []int
	res := 0
	for _, v := range divisors {
		if v == 0 {
			continue
		}
		for j := v; j < limit; j++ {
			if j % v == 0 && !slices.Contains(listMult, j){
				listMult = append(listMult, j)
				res += j
			}
		}
	}
	return res
	
}
