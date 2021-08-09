/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: heap_lt313.go
 * @Version: 1.0.0
 * @Data: 2021/8/9 6:50 PM
 */
package leetcode

import "container/heap"

type ugly []int

func nthSuperUglyNumber(n int, primes []int) int {
	ans := 1
	var data ugly
	cache := map[int]bool{}

	heap.Init(&data)
	heap.Push(&data, 1)
	cache[1] = true
	for i := 0; i < n; i++ {
		ans = heap.Pop(&data).(int)
		for _, v := range primes {
			if _, ok := cache[ans * v]; !ok {
				heap.Push(&data, ans * v)
				cache[ans * v] = true
			}
		}
	}
	return ans
}

func (h ugly) Len() int {
	return len(h)
}

func (h ugly) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h ugly) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *ugly) Push(n interface{})  {
	*h = append(*h, n.(int))
}

func (h *ugly) Pop() interface{} {
	t := *h
	r := t[len(t) - 1]
	*h = t[:len(t) - 1]
	return r
}