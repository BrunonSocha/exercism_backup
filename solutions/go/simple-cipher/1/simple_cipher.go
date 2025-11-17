package cipher

import (
	"math"
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift struct {
	values int
}

type vigenere struct {
	values []int
}


func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if math.Abs(float64(distance)) > 25 || math.Abs(float64(distance)) < 1 {
		return nil
	} 
	return shift{values: distance}
}

func (c shift) Encode(input string) string {
	res := ""
	input = strings.ToLower(input)
	for _, v := range input {
		if !unicode.IsLetter(v) {
			continue
		}
		k := int(v)
		k += c.values 
		if k > 122 {
			k = (k - 122) + 96
		} else if k < 97 {
			k = 122 - (96 - k)
		}
		res += string(byte(k))
	}
	return res 
}

func (c shift) Decode(input string) string {
	res := ""
	input = strings.ToLower(input)
	for _, v := range input {
		if !unicode.IsLetter(v) {
			continue
		}
		k := int(v)
		k -= c.values 
		if k < 97 {
			k = 122 - (96 - k)
		} else if k > 122 {
			k = (k - 122) + 96
		}
		res += string(byte(k))
	}
	return res
}

func NewVigenere(key string) Cipher {
	aCounter := 0
	var values []int
	if key == "" {
		return nil
	}
	for _, v := range key {
		if !unicode.IsLetter(v) || unicode.IsUpper(v) {
			return nil
		} else if v == 'a' {
			aCounter++
		}
		values = append(values, int(v)-97)
	}
	if aCounter == len(key) {
		return nil
	}
	return vigenere{values: values}
}

func (v vigenere) Encode(input string) string {
	res := ""
	i := 0
	for _, l := range input {
		if i == len(v.values) {
			i = 0
		}
		c := shift{values: v.values[i]}
		val := c.Encode(string(l))
		if val == "" {
			continue
		}
		res += val
		i++
	}
	return res
}

func (v vigenere) Decode(input string) string {
	res := ""
	i := 0
	for _, l := range input {
		if i == len(v.values) {
			i = 0
		}
		c := shift{values: v.values[i]}
		val := c.Decode(string(l))
		if val == "" {
			continue
		}
		res += val
		i++
	}
	return res
}
