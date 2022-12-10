package leetcode

import "fmt"

func MatchStr() {
	s := "abc"
	p := ".b*"

	res := isMatch(s, p)
	fmt.Println(res)
}

func isMatch(s string, p string) bool {
	for i := range s {
		if s[i] == p[i] || p[i] == '.' || p[i] == '*' {
			continue
		} else {
			return false
		}
	}
	return true
}
