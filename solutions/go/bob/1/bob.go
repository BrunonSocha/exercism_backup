// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}
	yell := false
	for _, l := range remark {
		if unicode.IsLetter(l){
			if unicode.IsUpper(l) {
				yell = true
			} else {
				yell = false
				break
			}
		}
	}
	var remarklist []rune
	remarklist = []rune(remark)
	question := remarklist[len(remarklist)-1] == '?'
	switch {
	case question && yell:
		return "Calm down, I know what I'm doing!"
	case question:
		return "Sure."
	case yell:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
