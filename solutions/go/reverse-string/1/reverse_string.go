package reverse

func Reverse(input string) string {
	var res []rune
	input2 := []rune(input)
	for i := len(input2)-1; i >= 0; i-- {
		res = append(res, input2[i])
	}
	return string(res)
}
