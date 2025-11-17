package prime

func Factors(n int64) []int64 {
	divisor := int64(2)
	var res []int64
	for n > 1 {
		if n % divisor == 0 {
			n /= divisor
			res = append(res, divisor)
			divisor = 2
		} else {
			divisor++
		}
	}
	return res
}
