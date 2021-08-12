/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_lt516.go
 * @Version: 1.0.0
 * @Data: 2021/8/12 11:19 PM
 */
package leetcode


func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}

	for i := range s {
		dp[i][i] = 1
		for j := i - 1; j >= 0; j-- {
			if s[j] == s[i] {
				dp[i][j] = dp[i- 1][j + 1] + 2
			} else {
				dp[i][j] = max(dp[i - 1][j], dp[i][j+1])
			}
		}
	}
	return dp[len(s) - 1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
