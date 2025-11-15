package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	sum := 0
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbnl := []rune(isbn)
	if len(isbnl) != 10 {
		return false
	}
	for i, v := range isbnl {
		val, err := strconv.Atoi(string(v))
		if err != nil {
			if v == 'X' && i == 9 {
				sum += 10*(10-i)
			} else {
				return false
			}
		}
		sum += val*(10-i)
	}
	return sum % 11 == 0
}
