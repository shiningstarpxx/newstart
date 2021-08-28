package leetcode

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	m := len(grid)
	dp := make([][]int, m)
	n := len(grid[0])
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = dp[0][i - 1] + grid[0][i]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
		}
	}

	fmt.Printf("%v\n", dp)
	return dp[m-1][n-1]
}
