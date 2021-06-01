/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  6:08 PM  31/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

import "errors"

type BinaryIndexedTree struct {
	tree []int
}

func NewBinaryIndexedTree(arr []int) *BinaryIndexedTree {
	n := len(arr)
	v := &BinaryIndexedTree{
		tree: make([]int, n + 1),
	}

	for i := 1; i <= n; i++ {
		v.tree[i] = 0;
	}

	for i := 0; i < n; i++ {
		v.UpdateBITree(i, arr[i])
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

/* return [begin, end] */
func (b* BinaryIndexedTree) GetRange(begin int, end int) (int, error) {
	if begin > end {
		return 0, errors.New("begin over end")
	}
	if begin == 0 {
		return b.GetSum(end)
	}
	l, e1 := b.GetSum(begin - 1)
	r, e2 := b.GetSum(end)

	if e1 != nil {
		return 0, e1
	} else if e2 != nil {
		return 0, e2
	} else {
		return r - l, nil
	}
}
