/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:57 PM  4/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "math"

func cherryPickup(grid [][]int) int {
	n := len(grid)
	tiles := make([][][]int, n)

	for i := range tiles {
		tiles[i] = make([][]int, n)
		for j := range tiles[i] {
			tiles[i][j] = make([]int, n)
			for k := range tiles[i][j] {
				tiles[i][j][k] = math.MinInt8
			}
		}
	}

	return int(math.Max(float64(dfsWithMemory(grid, 0, 0, 0, &tiles)), float64(0)))
}

func dfsWithMemory(grid [][]int, x1, y1, x2 int, c *[][][]int) int {
	// 边界情况
	y2 := x1 + y1 - x2
	if x1 >= len(grid) || y1 >= len(grid) || x2 >= len(grid) || y2 >= len(grid) ||
		grid[x1][y1] == -1 || grid[x2][y2] == -1 {
		return -1
	}
	// 合法终止
	if x1 == len(grid)-1 && y1 == len(grid)-1 {
		return grid[x1][y1]
	}
	// 已经访问过
	if (*c)[x1][y1][x2] != math.MinInt8 {
		return (*c)[x1][y1][x2]
	}

	res := int(math.Max(
		math.Max(float64(dfsWithMemory(grid, x1+1, y1, x2, c)), float64(dfsWithMemory(grid, x1, y1+1, x2, c))),
		math.Max(float64(dfsWithMemory(grid, x1+1, y1, x2+1, c)), float64(dfsWithMemory(grid, x1, y1+1, x2+1, c)))))
	// 这条路径不行
	if res < 0 {
		(*c)[x1][y1][x2] = -1
		return -1
	}

	res += grid[x1][y1]
	if x2 != x1 {
		res += grid[x2][y2]
	}

	(*c)[x1][y1][x2] = res
	return res
}
