/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:34 PM  24/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package example

import (
	"fmt"
	"math"
)

/* http://zxi.mytechroad.com/blog/dynamic-programming/leetcode-221-maximal-square/ */

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue;
			}
			dp[i][j] = 1
			if i != 0 && j != 0 {
				dp[i][j] = int(math.Min(float64(dp[i-1][j-1]), math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])))) + 1
			}
			res = int(math.Max(float64(res), float64(dp[i][j] * dp[i][j])))
			fmt.Printf("%v\n", dp)
		}
	}
	return res
}
