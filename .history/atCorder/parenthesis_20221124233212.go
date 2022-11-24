package atCorder

import "fmt"

func Parenthesis() {
	brackets := "(()())(())"

	res := parenthesis(brackets)
	fmt.Println(res)
}

func parenthesis(brackets string) bool {
	arr := []string{}

	for _, b := range brackets {
		switch string(b) {
		case "(" :
			arr = append(arr, string(b))
		case ")":
			arr = append(arr[0:len(arr)-2])
		}
	}

	fmt.Println(arr)

	return true
}