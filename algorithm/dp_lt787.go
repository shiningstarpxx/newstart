/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_lt787.go
 * @Version: 1.0.0
 * @Data: 2021/8/24 3:17 PM
 */
package algorithm

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][src] = 0  // Init

	for i := 1; i <= k + 1; i++ {
		for _, v := range flights {
			from, to, weight := v[0], v[1], v[2]
			dp[i][to] = int(math.Min(float64(dp[i][to]), float64(dp[i - 1][from] + weight)))  // Transition
		}
	}

	ans := math.MaxInt32
	for i := range dp {
		ans = int(math.Min(float64(dp[i][dst]), float64(ans)))
	}
	if ans == math.MaxInt32 {
		ans = -1
	}

	return ans
}
