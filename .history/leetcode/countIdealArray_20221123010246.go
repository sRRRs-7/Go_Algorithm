package leetcode

import "fmt"

func CountIdealArray() {
	num := 5
	maxValue := 3

	res := countIdealArray(num, maxValue)
	fmt.Printf("ideal array: %d \n", res)
}

func countIdealArray(num, maxValue int) int {
	res := 0

	res += 1 // constant res++

	for i := 1; i <= maxValue; i++ {
		if i + i > maxValue {
		} else if i + i < maxValue {
			res += (maxValue-1) * (num - 1)
			res -= 1
		}
	}

	return res
}
