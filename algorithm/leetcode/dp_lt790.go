/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:14 PM  28/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func numTilings(n int) int {
	// 2 * 1 : |        1
	// 2 * 2 : ||, =    2
	// 2 * 3 : |=, ||| =|   L7 [] 5
	dp := make([][3]int, n + 1)
	dp[1][0] = 1
	dp[1][2] = 0
	dp[1][1] = 0
	dp[2][0] = 2
	dp[2][1] = 1
	dp[2][2] = 1

	for i := 3; i <= n; i++ {
		dp[i][0] = dp[i - 2][0] + dp[i - 1][0] + dp[i - 1][1] + dp[i - 1][2]
		dp[i][0] %= 1e9 + 7
		dp[i][1] = dp[i - 2][0] + dp[i - 1][2]
		dp[i][2] = dp[i - 2][0] + dp[i - 1][1]
	}
	return dp[n][0]
}
