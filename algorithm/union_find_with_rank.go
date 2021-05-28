/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:50 PM  28/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

type UnionFindWithRank struct {
	parent []int
	count int
	ranks []int
}

// 这里需要声明的是：N个元素映射到了这个集合上，如果不是这样的集合，需要做一些换算变成是[0，n-1]可以表示的
func NewUnionFindWithRank(n int) *UnionFindWithRank {
	v := &UnionFindWithRank{
		parent: make([]int, n),
		count : n,
		ranks: make([]int, n),
	}

	for i := 0; i < n; i++ {
		v.parent[i] = i
		v.ranks[i] = 1
	}

	return v
}

func (uf *UnionFindWithRank) Find(x int) int {
	for ; uf.parent[x] != x; {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf *UnionFindWithRank) Union(x int, y int)  {
	rootx := uf.Find(x)
	rooty := uf.Find(y)

	if rooty == rootx {
		return
	}

	if uf.ranks[rootx] > uf.ranks[rooty] {
		uf.parent[rooty] = rootx
	} else if uf.ranks[rootx] < uf.ranks[rooty] {
		uf.parent[rootx] = rooty
	} else {
		uf.parent[rootx] = rooty
		uf.ranks[rooty]++
	}
	uf.count--
}

func (uf *UnionFindWithRank) IsConnected(x int, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFindWithRank) Count() int {
	return uf.count
}
