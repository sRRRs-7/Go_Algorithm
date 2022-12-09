package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func SubSequence() {
	nums := []int{5, -7, 3, 5}
	goal := 6

	res := subSequence(nums, goal)
	fmt.Println("sub sequence sum: ", res)
}

func subSequence(nums []int, goal int) int {
	mid := len(nums) / 2
	left := getAllSum(nums[:mid])
	right := getAllSum(nums[mid:])

	sort.Ints(left)
	sort.Ints(right)

	num := getMinVal(left, right, goal)
	return num
}

func getAllSum(nums []int) []int {
	n := len(nums)

	var res []int

	var iter func(idx, sumSoFar int)
	iter = func(idx, sumSoFar int) {
		if idx == n {
			res = append(res, sumSoFar)
			return
		}

		iter(idx+1, sumSoFar)
		iter(idx+1, sumSoFar+nums[idx])
	}

	iter(0, 0)

	return res
}

func getMinVal(left, right []int, goal int) int {
	minSoFar := math.MaxInt

	i := 0
	j := len(right) - 1

	for i < len(left) && j >= 0 {
		sum := left[i] + right[j]
		minSoFar = min(minSoFar, abs(goal, sum))

		if sum < goal {
			i++
		} else if sum > goal {
			j--
		} else {
			return 0
		}
	}

	return minSoFar
}
