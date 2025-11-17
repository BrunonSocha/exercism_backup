package armstrong

import (
	"strconv"
	"math"
)

func IsNumber(n int) bool {
	nStr := strconv.Itoa(n)
	res := 0
	for _, v := range nStr {
		val, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		res += int(math.Pow(float64(val), float64(len(nStr))))
	}
	return res == n
}
