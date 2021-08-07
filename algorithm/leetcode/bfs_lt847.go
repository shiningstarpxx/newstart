/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: bfs_lt847.go
 * @Version: 1.0.0
 * @Data: 2021/8/6 9:11 AM
 */
package leetcode

type pathNode struct {
	current, visitPath, distance int
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	q := make([]pathNode, n)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, 1 << n)
		visited[i][1<<i] = true
		q[i] = pathNode{i, 1<<i, 0}
	}

	for len(q) > 0 {
		top := q[0]
		if top.visitPath == (1 << n - 1) {
			return top.distance
		}
		q = q[1:]
		for _, v := range graph[top.current] {
			mask := top.visitPath | (1 << v)
			if !visited[v][mask] {
				visited[v][mask] = true
				q = append(q, pathNode{v, mask, top.distance + 1})
			}
		}
	}
	return 0
}
