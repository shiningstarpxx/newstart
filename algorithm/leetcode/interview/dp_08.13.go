/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_08.13.go
 * @Version: 1.0.0
 * @Data: 2021/9/4 10:24 AM
 */
package interview

import (
	"math"
	"sort"
)

func pileBox(box [][]int) int {
	box = sortBox(box)

	ans := 0
	n := len(box)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = box[i][2]
		for j := 0; j < i; j++ {
			if box[j][0] < box[i][0] && box[j][1] < box[i][1] && box[j][2] < box[i][2] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j] + box[i][2])))
			}
		}
		ans = int(math.Max(float64(ans), float64(dp[i])))
	}
	return ans
}

func sortBox(box [][]int) [][]int {
	sort.Slice(box, func(i, j int) bool {
		if box[i][0] != box[j][0] {
			return box[i][0] < box[j][0]
		}
		if box[i][1] != box[j][1] {
			return box[i][1] < box[j][1]
		}
		return box[i][2] < box[j][2]
	})
	return box
}
