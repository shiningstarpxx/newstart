/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_lt600.go
 * @Version: 1.0.0
 * @Data: 2021/9/11 8:37 PM
 */
package leetcode

func findIntegers(n int) int {
	dp := make([]int, 32)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < 31; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}

	res, pre := 0, 0
	for i := 29; i >= 0; i-- {
		if n & (1 << i) != 0 {

			res += dp[i + 1]
			n -= (1<<i)
			if pre == 1 {
				break
			}
			pre = 1
		} else {
			pre = 0
		}

		if i == 0 {
			res++
		}
	}

	return res
}
