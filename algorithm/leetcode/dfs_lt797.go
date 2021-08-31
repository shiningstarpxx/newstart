/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dfs_lt797.go
 * @Version: 1.0.0
 * @Data: 2021/8/25 7:41 PM
 */
package leetcode

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	var ans [][]int
	stk := []int{0}
	var dfs func(int)
	dfs = func(i int) {
		if i == n - 1 {
			ans = append(ans, append([]int{}, stk...))
			return
		}
		for _, v := range graph[i] {
			stk = append(stk, v)
			dfs(v)
			stk = stk[:len(stk) - 1]
		}
	}
	dfs(0)
	return ans
}
