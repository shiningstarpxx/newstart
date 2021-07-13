/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:41 PM  13/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func swimInWater(grid [][]int) int {
	n := len(grid)
	index := make([]int, n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			index[grid[i][j]] = calcIndex(i, j , n)
		}
	}

	//fmt.Printf("%v\n", index)
	uf := NewUnionFind778(n * n)
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i:= 0; i < n * n; i++ {
		x, y := index[i] / n, index[i] % n
		for _, d := range dir {
			newX, newY := x + d[0], y + d[1]
			if inArea(newX, newY, n) && grid[newX][newY] <= i {
				//fmt.Printf("union %d, %d\n", index[i], calcIndex(newX, newY, n))
				uf.Union(index[i], calcIndex(newX, newY, n))
			}
			if uf.IsConnected(0, n * n - 1) {
				return i
			}
		}
	}
	return -1
}

func calcIndex(x, y, n int) int {
	return x * n + y
}

func inArea(x, y, n int) bool {
	if x < 0 || x >= n || y < 0 || y >= n {
		return false
	}
	return true
}

type UnionFind778 struct {
	parent []int
	count int
}

// 这里需要声明的是：N个元素映射到了这个集合上，如果不是这样的集合，需要做一些换算变成是[0，n-1]可以表示的
func NewUnionFind778(n int) *UnionFind778 {
	v := &UnionFind778{
		parent: make([]int, n),
		count : n,
	}

	for i := 0; i < n; i++ {
		v.parent[i] = i
	}

	return v
}

func (uf *UnionFind778) Find(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf *UnionFind778) Union(x int, y int)  {
	rootx := uf.Find(x)
	rooty := uf.Find(y)

	if rooty == rootx {
		return
	}

	uf.parent[rootx] = rooty
	uf.count--
}

func (uf *UnionFind778) IsConnected(x int, y int) bool {
	return uf.Find(x) == uf.Find(y)
}