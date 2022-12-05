package regex

import (
	"fmt"
	"regexp"
)

func RegexFunc() {
	fmt.Println(checkRegex("a*c", "aaaeccc"))
}

func checkRegex(reg, s string) bool {
	return regexp.MustCompile(reg).MatchString(s)
}