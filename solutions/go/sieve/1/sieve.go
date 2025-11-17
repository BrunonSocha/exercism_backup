package sieve

func Sieve(limit int) []int {
	var res []int
	if limit >= 2 {
		res = append(res, 2)
	}
	for j := 3; j <= limit; j++ {
		prime := true
		for _, k := range res {
			if j % k == 0 {
				prime = false
			}
		}
		if prime {
			res = append(res, j)
		}
	}
	return res
}
