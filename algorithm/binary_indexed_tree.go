/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  6:08 PM  31/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

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
		v.UpdateBITree(n, i, arr[i])
	}
	return v
}

func (b* BinaryIndexedTree) UpdateBITree(size int, index int, v int) {
	index = index + 1
	for ; index <= size; {
		b.tree[index] += v
		index += index & (-index)
	}
}

