/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_dijkstras_lt743.go
 * @Version: 1.0.0
 * @Data: 2021/8/2 5:29 PM
 */
package leetcode

import "math"

func networkDelayTime(times [][]int, n int, k int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt32 / 2
		}
	}
	for _, data := range times{
		x := data[0] - 1
		y := data[1] - 1
		g[x][y] = data[2]
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32 / 2
	}
	dist[k - 1] = 0
	used := make([]bool, n)

	for i := 0; i < n; i++ {
		c := -1
		for j := range used {
			if !used[j] && (c == -1 || dist[j] < dist[c]) {
				c = j
			}
		}
		used[c] = true

		for k := 0; k < n; k++ {
			dist[k] = int(math.Min(float64(dist[k]), float64(dist[c] + g[c][k])))
		}
	}

	ans := 0
	for _, d := range dist {
		ans = int(math.Max(float64(ans), float64(d)))
	}
	if ans == math.MaxInt32 / 2 {
		return -1
	}
	return ans
}
