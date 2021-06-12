/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  5:51 PM  6/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

import "errors"

type STree struct {
	tree []int
	size int
}

func NewSTree(arr []int) *STree {
	l := len(arr)
	v := &STree{
		tree : make([]int, l * 2),
		size : l,
	}

	// insert leaf node in tree
	for i := 0; i < l; i++ {
		v.tree[l + i] = arr[i]
	}

	// build the tree by calculating parents
	for i := l - 1; i > 0; i-- {
		v.tree[i] = v.tree[i << 1] + v.tree[i << 1 | 1]
	}
	return v
}

// Update Node value by index
func (s* STree) UpdateTreeNode(index, value int) {
	index = index + s.size
	s.tree[index] = value

	for i := index; i > 1; i >>= 1 {
		s.tree[i >> 1] = s.tree[i] + s.tree[i ^ 1];
	}
}

// Calc Sum in range [l, r) by index
func (s* STree) Query(l, r int) (int, error) {
	if l < 0 || r < 0 || l >= s.size || r >= s.size {
		return 0, errors.New("out of Range")
	}

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
	return sum, nil
}