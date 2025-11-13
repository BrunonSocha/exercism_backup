package parsinglogfiles

import (
	"regexp"
	"fmt"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`\A[\[](TRC|DBG|INF|WRN|ERR|FTL)[\]].+`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`[<](~|=|\*|-)*[>]`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	res := 0
	re := regexp.MustCompile(`["](?i).*password.*["]`)
	for _, v := range lines {
		matches := re.FindAllString(v, -1)
		res += len(matches)
	}
	return res
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(end-of-line)[0-9]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var res []string
	reUsrN := regexp.MustCompile(`(User)\s+\S*`)
	reUsr := regexp.MustCompile(`(User)(\s+)`)
	tag := "[USR]"
	for _, v := range lines {
		if reUsrN.MatchString(v) {
			usrAndTag := reUsrN.FindString(v)
			nick := reUsr.ReplaceAllString(usrAndTag, "")
			res = append(res, fmt.Sprintf("%s %s %s", tag, nick, v))
		} else {
			res = append(res, v)
		}
	}
	return res
}
