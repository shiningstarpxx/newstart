/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: bfs_topological_lt802.go
 * @Version: 1.0.0
 * @Data: 2021/8/5 11:56 AM
 */
package leetcode

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	nodeCount := make([]int, n)
	var q []int
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		for _, v := range graph[i] {
			cache[v] = append(cache[v], i)
		}
		nodeCount[i] = len(graph[i])
		if nodeCount[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		d := q[len(q) - 1]
		q = q[:len(q) - 1]
		for _, v := range cache[d] {
			nodeCount[v]--
			if nodeCount[v] == 0 {
				q = append(q, v)
			}
		}
	}

	ans := []int{}
	for i := 0; i < n; i++ {
		if nodeCount[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
