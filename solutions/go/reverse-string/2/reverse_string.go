package reverse

func Reverse(input string) string {
	var res string
	for _, v := range input {
		res = string(v) + res
	}
	return string(res)
}
