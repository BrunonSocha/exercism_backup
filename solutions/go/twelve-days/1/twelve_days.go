package twelve
import "fmt"
func Verse(i int) string {
	days := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	items := []string{"a Partridge in a Pear Tree.", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}
	baseStr := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", days[i-1])
	if i == 1 {
		return baseStr + " " + items[i-1]  
	} else {
		resStr := ""
		for v := i; v > 1; v-- {
			resStr += " " + items[v-1] + ","
		}
		return baseStr + resStr + " and " + items[0]
	}
}

func Song() string {
	res := "" 
	for i := 1; i < 13; i++ {
		if i == 12 {
			res += fmt.Sprint(Verse(i))
			break
		}
		res += fmt.Sprintln(Verse(i))
	}
	return res
}
