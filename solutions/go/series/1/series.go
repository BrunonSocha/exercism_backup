package series

func All(n int, s string) []string {
	var res []string
	for i := 0; i < len(s)-(n-1); i++ {
		res = append(res, s[i:(i+n)])
	} 
	return res
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

//func UnsafeFirst(n int, s string) (first string, ok bool) {
//	return s[0:n], len(s) <= n
//}
