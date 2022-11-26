package atCorder

import "fmt"

func GridMath() {
	grid := [][]int{
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
}

	res := gridMath(grid)
	fmt.Println("grid sum: ", res)
}

func gridMath(grid [][]int) [][]int {
	sumGrid := grid

	for i, r := range grid {
		var sum = 0
		for ii := range r {
			for idx, y := range r {
				sum += y
				if idx != ii { sum += grid[idx][ii] }
			}
			sumGrid[i][ii] = sum
		}
		sum = 0
	}

	return sumGrid
}