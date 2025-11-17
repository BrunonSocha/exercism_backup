package phonenumber

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	res := ""
	for _, v := range phoneNumber {
		if unicode.IsDigit(v) {
			res += string(v)
		} 
	}
	if len(res) == 11 && res[0] == '1' {
		res = res[1:]
	}
	if len(res) != 10 {
		return "", errors.New("Wrong number length.")
	}
	n1, err := strconv.Atoi(string(res[0]))
	if err != nil {
		return "", err
	}
	n2, err := strconv.Atoi(string(res[3]))
	if err != nil {
		return "", err
	}
	if n1 < 2 || n2 < 2 {
		return "", errors.New("Invalid N numbers.")
	}
	return res, nil 
}

func AreaCode(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return cleanNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	areaCode, err := AreaCode(cleanNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", areaCode, cleanNumber[3:6], cleanNumber[6:]), nil
}
