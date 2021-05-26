/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:06 AM  26/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description: Union Find Or Disjoint Set
*/
package algorithm

type UnionFind struct {
	parent []int
	count int
}

// 这里需要声明的是：N个元素映射到了这个集合上，如果不是这样的集合，需要做一些换算变成是[0，n-1]可以表示的
func NewUnionFind(n int) *UnionFind {
	v := &UnionFind{
		parent: make([]int, n),
		count : n,
	}

	for i := 0; i < n; i++ {
		v.parent[i] = i
	}

	return v
}

func (uf *UnionFind) Find(x int) int {
	for ; uf.parent[x] != x; {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf *UnionFind) Union(x int, y int)  {
	rootx := uf.Find(x)
	rooty := uf.Find(y)

	if rooty == rootx {
		return
	}

	uf.parent[rootx] = rooty
	uf.count--
}

func (uf *UnionFind) IsConnected(x int, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Count() int {
	return uf.count
}