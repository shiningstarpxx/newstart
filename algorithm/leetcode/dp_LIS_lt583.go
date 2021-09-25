/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_LIS_lt583.go
 * @Version: 1.0.0
 * @Data: 2021/9/25 11:26 AM
 */
package leetcode

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m + 1)
	for i := range dp {
		dp[i] = make([]int, n + 1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max583(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return m + n - 2 * dp[m][n]
}

func max583(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
