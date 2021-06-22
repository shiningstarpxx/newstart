/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  5:48 PM  20/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package example

import (
	"container/heap"
	"sort"
)

type sortable [][]int

func (s sortable) Len() int {
	return len(s)
}
func (s sortable) Less(i,j int) bool {
	if s[i][0]==s[j][0] {
		return s[i][1]<=s[j][1]
	}
	return s[i][0]<=s[j][0]
}
func (s sortable) Swap(i,j int) {
	s[i],s[j]=s[j],s[i]
}

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Erase(x interface{}) {
	c := make([]int, 1)
	if len(*h) > 0 && (*h)[len(*h) - 1] != x.(int) {
		c = append(c, h.Pop().(int))
	}
	h.Pop()
	for _, v := range c {
		h.Push(v)
	}
}

func (h *IntHeap) Top() interface{} {
	return (*h)[len(*h) - 1]
}

func getSkyline(buildings [][]int) [][]int {
	d := make([][]int, 2 * len(buildings))
	index := 0
	for _, v := range buildings {
		d[index] = make([]int, 2)
		d = append(d, []int{v[0], -v[2]})
		index++
		d[index] = make([]int, 2)
		d = append(d, []int{v[1], -v[2]})
		index++
	}
	sort.Sort(sortable(d))

	last := []int{0, 0}
	h := &IntHeap{}
	heap.Init(h)
	r := make([][]int, 0)
	for _, v := range d {
		if v[1] < 0 {
			h.Push(-v[1])
		} else {
			h.Erase(v[1])
		}

		max := h.Top()
		if max.(int) != last[1] {
			last[0] = v[0]
			last[1] = max.(int)
			r = append(r, last)
		}
	}
	return r
}

