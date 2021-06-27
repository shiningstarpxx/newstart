/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  5:33 PM  15/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package example

import (
	"errors"
	"sort"
)

/* https://zxi.mytechroad.com/blog/algorithms/array/leetcode-315-count-of-smaller-numbers-after-self/ */

func countSmallerWithBIT(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	cache := make(map[int]int, len(nums))
	for i, v := range sorted {
		cache[v] = i + 1
	}

	bit := NewBinaryIndexedTreeWithSize(len(nums))
	reverseSlice(nums)

	res := make([]int, len(nums))
	for i, v := range nums {
		// find smaller than this one and count sum
		sum, _ := bit.GetSum(cache[v] - 1)
		res[i] += sum
		// update current value sum
		bit.UpdateBITree(cache[v], 1)
	}
	reverseSlice(res)
	return res
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

type BinaryIndexedTree struct {
	tree []int
}

func NewBinaryIndexedTreeWithSize(size int) *BinaryIndexedTree  {
	v := &BinaryIndexedTree{
		tree: make([]int, size + 1),
	}
	return v
}

func (b* BinaryIndexedTree) UpdateBITree(index int, v int) {
	index = index + 1
	for ; index < len(b.tree); {
		b.tree[index] += v
		index += index & (-index)
	}
}

func (b* BinaryIndexedTree) GetSum(index int) (int, error){
	if index >= len(b.tree) - 1 || index < 0 {
		return 0, errors.New("out of tree range")
	}
	sum := 0
	index = index + 1
	for ;index > 0; {
		sum += b.tree[index]
		index = index - (index & (-index))
	}
	return sum, nil
}