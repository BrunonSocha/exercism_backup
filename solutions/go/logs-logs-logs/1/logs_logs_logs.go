package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	characters := map[rune]string {'❗': "recommendation", '🔍': "search", '☀': "weather"}
	for _, v := range log {
		for u, w := range characters {
			if v == u {
				return w 
			}
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newString := ""
	for _, v := range log {
			
		if v == oldRune {
			newString += string(newRune)
			continue
		}
		newString += string(v)
	}
	return newString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit	
}
