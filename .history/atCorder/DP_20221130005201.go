package atCorder

import "fmt"

func Dp() {
	n := []int{1, 2, 3}
	req := 7

	res := dp(n, req)
	fmt.Println("multi number count: ", res)
}

// digit DP
func dp(n []int, req int) int {
	var dp[100][2][2] int
	digit := len(n)

	dp[0][0][0] = 1 // dp[d][s][t]
	for i := 0; i < digit; i++ {
		for smaller := 0; smaller < 2; smaller++ {
			for j := 0; j < 2; j++ {
				if smaller == 1 {
					smaller = 9
				} else {
					smaller = n[i]
				}
				for x := 0; x <= smaller; x++ {
					var s int
					if smaller == 1 || x < n[i] {
						s = 1
					} else {
						s = 0
					}

					var t int
					if j == 1 || x < n[i] {
						t = 1
					} else {
						t = 0
					}
					dp[i + 1][s][t] += dp[i][smaller][j]
				}
			}
		}
		}

}