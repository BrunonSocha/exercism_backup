package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	for _, v := range pt {
		if !unicode.IsDigit(v) && !unicode.IsLetter(v) {
			pt = strings.ReplaceAll(pt, string(v), "")
		}
	}
	c := 1
	r := 0
	alternate := false
	if math.Sqrt(float64(len(pt))) == math.Trunc(math.Sqrt(float64(len(pt)))) {
		c = int(math.Sqrt(float64(len(pt))))
		r = c
	} else {
		for r * c < len(pt) && c >= r && c-r <= 1 {
			if alternate {
				c++
				alternate = false
			} else {
				r++
				alternate = true
			}
		}
	}
	var ptSli []string
	buffer := ""
	for _, v := range strings.ToLower(pt) {
		if len(buffer) == c {
			ptSli = append(ptSli, buffer)
			buffer = ""
		}
		buffer += string(v)
	}
	if len(buffer) > 0 {
		ptSli = append(ptSli, buffer + strings.Repeat(" ", c-(len(buffer))))
	}
	var transPt []string
	buffer = ""
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			buffer += string(ptSli[j][i])
		}
		if len(buffer) == r {
			transPt = append(transPt, buffer)
			buffer = ""
		}
	}

	res := ""
	for _, v := range transPt {
		res += (v + " ")
	}
	if res == "" {
		return res
	}
	return res[:(len(res)-1)] 
}
