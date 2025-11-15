package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	resStr := ""
	bands := map[string]string{"black": "0", "brown": "1", "red": "2", "orange": "3", "yellow": "4", "green": "5", "blue": "6", "violet": "7", "grey": "8", "white": "9"}
	for _, v := range colors[:2] {
		resStr += bands[v]
	}
	res, err := strconv.Atoi(resStr)
	if err != nil {
		panic(err)
	}
	return res
}
