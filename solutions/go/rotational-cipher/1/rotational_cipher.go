package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	var res []byte
	for _, v := range []byte(plain) {
		if v >= 'a' && v <= 'z'{
			res = append(res, byte((((int(v) + shiftKey) - 97) % 26) + 97))
		} else if v >= 'A' && v <= 'Z' {
			res = append(res, byte((((int(v) + shiftKey) - 65) % 26) + 65))
		} else {
			res = append(res, v)
		}
	}
	
	return string(res)
}
