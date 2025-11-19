package wordy

import (
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Answer(question string) (int, bool) {
	operations := []string{"plus", "minus", "divided", "multiplied"}
	var operands []int
	var opsExec []string
	questionSli := strings.Split(question, " ")
	opsInRow := 0
	opdsinRow := 0
	for _, v := range questionSli {
		
		v = strings.TrimSuffix(v, "?")
		if slices.Contains(operations, v) {
			opsExec = append(opsExec, v) 
			opsInRow++
			opdsinRow = 0

		}
		vInt, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		operands = append(operands, vInt)
		opdsinRow++
		opsInRow = 0
		if opdsinRow == 2 || opsInRow == 2 {
			return 0, false
		}
	}
	if len(operands) == 1 && len(opsExec) == 0 && unicode.IsDigit(rune(question[len(question)-2])) {
		return operands[0], true
	} else if operands == nil || opsExec == nil {
		return 0, false
	}
	if len(operands) != len(opsExec)+1 {
		return 0, false
	}
	res := operands[0]
	for i := 1; i < len(operands); i++ {
		switch opsExec[i-1] {
		case "plus":
			res += operands[i]
		case "minus":
			res -= operands[i]
		case "divided":
			res /= operands[i]
		case "multiplied":
			res *= operands[i]
		}
	}
	return res, true
}
