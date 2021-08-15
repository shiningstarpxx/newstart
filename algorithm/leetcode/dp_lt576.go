/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_lt576.go
 * @Version: 1.0.0
 * @Data: 2021/8/15 5:49 PM
 */
package leetcode

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, maxMove + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	const mod int = 1e9 + 7
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dp[0][startRow][startColumn] = 1
	ans := 0
	for k := 0; k < maxMove; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dp[k][i][j] > 0 {
					for _, v := range dir {
						x := i + v[0]
						y := j + v[1]
						if x >= 0 && x < m && y >= 0 && y < n {
							dp[k + 1][x][y] = (dp[k + 1][x][y] + dp[k][i][j]) % mod
						} else {
							ans = (ans + dp[k][i][j]) % mod
						}
					}
				}
			}
		}
	}
	return ans
}
