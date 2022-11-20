package utils

import (
	"fmt"
)

func BaseNumberConvert() {
	num := 10
	base := 2

	res := baseNumberConvert(num, base)
	fmt.Println("convert base number", res)
}

func baseNumberConvert(num, base int) int {
	remainder := 0
	quotient := 0
	result := []int{}

	for {
		remainder = num % base
		quotient = num / base
		num = quotient
		result = append(result, remainder)
		if quotient == 1 {
			result = append(result, quotient)
			break
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	for _, n ;= range result {}

	return res
}