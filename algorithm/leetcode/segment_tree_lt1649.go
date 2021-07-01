/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:41 AM  19/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"fmt"
	"sort"
)

const (
	STNumRange = 100000 + 1
	StMod = 1e9+7
)

// CreateSortedArrayST
/*
leetcode : [1 5 6 2]
1. create sorted array [1 2 5 6]
2. create segment tree from sorted array, but only for index [1->0, 2->1, 5->2, 6->3]
3. for range primary array [1 5 6 2],  index is current pivot
	3.1 calc current smaller = sum[0, index - 1], bigger = sum (index + 1, len(array) - 1)
	3.2 min(smaller, bigger)
    3.3 add(index, 1)
 */
func CreateSortedArrayST(instructions []int) int {
	sorted := make([]int, len(instructions))
	copy(sorted, instructions)
	sort.Ints(sorted)

	cache := make(map[int]int)
	index := 0
	for _, v := range sorted {
		if _, ok := cache[v]; !ok {
			cache[v] = index
			index++
		}
	}

	ret := 0
	tree := NewSTreeBySize(index)
	for _, v := range instructions {
		// tree.Print()
		pivot := cache[v]
		smaller := tree.Query(0, pivot)
		bigger := tree.Query(pivot + 1, index)
		fmt.Printf("%d, %d\n", smaller, bigger)
		ret += min1649(smaller, bigger)
		ret %= StMod
		tree.UpdateTreeNode(pivot, 1)
	}
	return ret
}

func min1649(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type STree struct {
	tree []int
	size int
}

// NewSTreeBySize Init an empty segment tree
func NewSTreeBySize(size int) *STree {
	v := &STree{
		tree : make([]int, size * 2),
		size : size,
	}
	for i := 0; i < 2 * size; i++ {
		v.tree[i] = 0
	}
	fmt.Printf("%v\n", v.tree)
	return v
}

// UpdateTreeNode Update Node value by index
func (s* STree) UpdateTreeNode(index, value int) {
	index = index + s.size
	s.tree[index] += value

	for i := index; i > 1; i >>= 1 {
		s.tree[i >> 1] = s.tree[i] + s.tree[i ^ 1]
	}
}

// Query Calc Sum in range [l, r) by index
func (s* STree) Query(l, r int) int {
	sum := 0
	l += s.size
	r += s.size
	for l < r {
		if l & 1 == 1 {
			sum += s.tree[l]
			l++
		}
		if r & 1 == 1 {
			r--
			sum += s.tree[r]
		}
		l >>= 1
		r >>= 1
	}
	return sum
}

func (s* STree) Print() {
	fmt.Printf("%v, %d\n", s.tree, len(s.tree))
}