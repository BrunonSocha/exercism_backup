package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	var digitsToDouble []int
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}
	idr := []rune(id)
	if len(idr) % 2 == 0 {
		for i := 0; i < len(idr); i++ {
			if i % 2 == 0 {
				num, err := strconv.Atoi(string(idr[i]))
				if err != nil {
					return false
				}
				num *= 2
				if num > 9 {
					num -= 9
				}
				digitsToDouble = append(digitsToDouble, num)
			} else {
				num, err := strconv.Atoi(string(idr[i]))
				if err != nil {
					return false
				}
				digitsToDouble = append(digitsToDouble, num)
			}
		}
	} else {
		for i := 0; i < len(idr); i++ {
			if i % 2 != 0 {
				num, err := strconv.Atoi(string(idr[i]))
				if err != nil {
					return false
				}
				num*=2
				if num > 9 {
					num -= 9
				}
				digitsToDouble = append(digitsToDouble, num)
			} else {
				num, err := strconv.Atoi(string(idr[i]))
				if err != nil {
					return false
				}
				digitsToDouble = append(digitsToDouble, num)
			}
		}
	}
	res := 0
	for i := 0; i < len(digitsToDouble); i++ {
		res += digitsToDouble[i]
	}
	if res % 10 == 0 {
		return true
	}
	return false
}
