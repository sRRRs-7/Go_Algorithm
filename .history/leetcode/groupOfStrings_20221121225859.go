package leetcode

import (
	"fmt"
	"strings"
)

func GroupOfStrings() {
	words := []string{"a","b","ab","cde"}
	res := groupOfStrings(words)
	fmt.Println(res)
}

func groupOfStrings(words []string) [2]int {
	res := [2]int{0, 0} // [divide number, max length group]
	maxLen := 0
	div := 0

	for i, w1 := range words {
		for _, w2 := range words {
			if strings.Contains(w1, w2) { maxLen += 1 }
		}
		if maxLen > 1 { div += 1 }
		if res[1] < maxLen { res[1] = maxLen }
		if i == len(words) { res[0] = len(words)-maxLen+1 }
		maxLen = 0
	}

	return res
}