/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: heap_17.14.go
 * @Version: 1.0.0
 * @Data: 2021/9/3 9:41 PM
 */
package interview

import (
	"container/heap"
	"sort"
)

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}

	h := &hp{arr[:k]}
	heap.Init(h)

	for _, v := range arr[k:] {
		if v < h.IntSlice[0] {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }