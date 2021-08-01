/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: heap_lt1337.go
 * @Version: 1.0.0
 * @Data: 2021/8/1 9:11 PM
 */
package leetcode

import (
	"container/heap"
	"sort"
)

type node struct {
	score int
	index int
}

type nodes []node

func kWeakestRows(mat [][]int, k int) []int {
	var data nodes
	heap.Init(&data)

	for i, j := range mat {
		score := sort.Search(len(j), func(index int) bool {
			return j[index] == 0
		})
		heap.Push(&data, node{score, i})
	}

	var ans []int
	for i := 0; i < k; i++ {
		v := heap.Pop(&data)
		ans = append(ans, v.(node).index)
	}
	return ans
}

func (h nodes) Len() int {
	return len(h)
}

func (h nodes) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h nodes) Less(i, j int) bool {
	return h[i].score < h[j].score || (h[i].score == h[j].score && h[i].index < h[j].index)
}

func (h *nodes) Push(n interface{})  {
	*h = append(*h, n.(node))
}

func (h *nodes) Pop() interface{} {
	t := *h
	r := t[len(t) - 1]
	*h = t[:len(t) - 1]
	return r
}