/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:24 PM  18/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package example

const (
	NumRange = 100000 + 1
	Mod = 1e9+7
)

func createSortedArrayBIT(instructions []int) int {
	tree := NewBinaryIndexedTreeWithSize1649(NumRange)
	res := 0

	for i, v := range instructions {
		// sum(<=v-1), current we have i - 1 + 1 = i numbers, so sum(>v)
		res += min(tree.GetSum(v - 1), i - tree.GetSum(v))
		res %= Mod
		// then input ithe number
		tree.UpdateBITree(v, 1)
	}
	return res
}

type BinaryIndexedTree1649 struct {
	tree []int
}

func NewBinaryIndexedTreeWithSize1649(size int) *BinaryIndexedTree1649  {
	v := &BinaryIndexedTree1649{
		tree: make([]int, size + 1),
	}
	return v
}

func (b* BinaryIndexedTree1649) UpdateBITree(index int, v int) {
	index = index + 1
	for ; index < len(b.tree); {
		b.tree[index] += v
		index += index & (-index)
	}
}

func (b* BinaryIndexedTree1649) GetSum(index int) (int){
	sum := 0
	index = index + 1
	for ;index > 0; {
		sum += b.tree[index]
		index = index - (index & (-index))
	}
	return sum
}