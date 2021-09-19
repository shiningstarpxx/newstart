/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_lt650.go
 * @Version: 1.0.0
 * @Data: 2021/9/19 3:28 PM
 */
package leetcode

import "math"

func minSteps(n int) int {
	dp := make([]int, n + 1)

	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j * j <= i; j++ {
			if i % j == 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[j] + i / j)))
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i / j] + j)))
			}
		}
	}
	return dp[n]
}