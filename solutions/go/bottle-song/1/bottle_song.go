package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	var res []string
	numbers := map[int]string{10:"Ten", 9:"Nine", 8:"Eight", 7:"Seven", 6:"Six", 5:"Five", 4:"Four", 3:"Three", 2:"Two", 1:"One", 0:"no"}
	bottle := "bottles"
	for i := startBottles; i > 0; i-- {
		if i == 1 {
			bottle = "bottle"
		}
		res = append(res, fmt.Sprintf("%s green %s hanging on the wall,", numbers[i], bottle))
		res = append(res, fmt.Sprintf("%s green %s hanging on the wall,", numbers[i], bottle))
		res = append(res, fmt.Sprint("And if one green bottle should accidentally fall,"))
		if i == 2 {
			bottle = "bottle"
		} else if i == 1 {
			bottle = "bottles"
		}
		res = append(res, fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(numbers[i-1]), bottle))
		if (startBottles - (i-1)) == takeDown {
			return res
		}
		if i != 1 {
			res = append(res, "")
		}
	}
	return res
}
