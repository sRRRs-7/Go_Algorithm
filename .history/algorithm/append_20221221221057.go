package algorithm

import "fmt"

func AddArray() {
	addArray()
}

func addArray() {
	arr := make([]int, 0)

	arr = append(arr, 1)
	fmt.Println(arr)

	arr = arr[:len(arr)]
	fmt.Println(arr)
}
