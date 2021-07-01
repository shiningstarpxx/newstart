/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  5:22 PM  5/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "sort"


func reversePairs(nums []int) int {
	n := len(nums)

	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	t := NewBIT(n)
	res := 0
	for _, v := range nums {
		index := upperBound(2 * int64(v), sorted)
		res += t.GetSum(n - index)
		index = upperBound(int64(v)-1, sorted)
		t.Update(n - index)
	}
	return res
}

func upperBound(key int64, num []int) int {
	left, right := 0, len(num)
	for ; left < right; {
		mid := left + (right - left) / 2
		if int64(num[mid]) > key {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

type BIT struct {
	tree []int
}

func NewBIT(n int) *BIT {
	v := &BIT{
		tree: make([]int, n+1),
	}
	for i := 1; i <= n; i++ {
		v.tree[i] = 0
	}
	return v
}

func (b *BIT) GetSum(i int) int {
	sum := 0
	for ; i > 0; {
		sum += b.tree[i]
		i -= i & (-i)
	}
	return sum
}

func (b *BIT) Update(i int) {
	for ; i < len(b.tree); {
		b.tree[i] += 1
		i += i & (-i)
	}
}