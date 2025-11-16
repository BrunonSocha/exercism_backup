package lsproduct

import (
	"strconv"
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var buffer []int
	var res int64  
	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	} else if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	for i, _ := range digits {
		var canRes int64 = 1 
		if len(digits) - i < span {
			break
		}
		for n := i; n < i+span; n++ {
			val, err := strconv.Atoi(string(digits[n]))
			if err != nil {
				return 0, err
			}
			buffer = append(buffer, val)
		}
		for _, v := range buffer {
			canRes *= int64(v)
		}
		if canRes > res || i == 0  {
			res = canRes
		}
		buffer = nil
	}
	return res, nil
}
