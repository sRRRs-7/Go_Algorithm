package leetcode

func CountIdealArray() {
	num := 2
	maxValue := 5

	countIdealArray(num, maxValue)
}

func countIdealArray(num, maxValue int) int {
	arr := []int{}
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}