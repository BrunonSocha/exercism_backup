package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	res := ""
	if input > 3999 || input <= 0 {
		return "", errors.New("Incorrect number.")
	}
	for {
		if input >= 1000 {
			res += "M"
			input -= 1000
		} else if input >= 900 {
			res += "CM"
			input -= 900
		} else if input >= 500 {
			res += "D"
			input -= 500
		} else if input >= 400 {
			res += "CD"
			input -= 400
		} else if input >= 100 {
			res += "C"
			input -= 100
		} else if input >= 90 {
			res += "XC"
			input -= 90
		} else if input >= 50 {
			res += "L"
			input -= 50
		} else if input >= 40 {
			res += "XL"
			input -= 40
		} else if input >= 10 {
			res += "X"
			input -= 10
		} else if input >= 9 {
			res += "IX"
			input -= 9
		} else if input >= 5 {
			res += "V"
			input -= 5
		} else if input >= 4 {
			res += "IV"
			input -= 4
		} else if input >= 1 {
			res += "I"
			input -= 1
		} else {
			break
		}
	}
	return res, nil
}
