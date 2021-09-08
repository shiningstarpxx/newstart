/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: heap_greedy_lt502.go
 * @Version: 1.0.0
 * @Data: 2021/9/8 10:15 PM
 */
package leetcode

import (
	"container/heap"
	"sort"
)

type pair struct {
	c int
	p int
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	data := make([]pair, n)
	for i := 0; i < n; i++ {
		data[i].p = profits[i]
		data[i].c = capital[i]
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].c < data[j].c
	})

	h := &hp{}
	for cur := 0; k > 0; k-- {
		for cur < n && data[cur].c <= w {
			heap.Push(h, data[cur].p)
			cur++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
	}
	return w
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }