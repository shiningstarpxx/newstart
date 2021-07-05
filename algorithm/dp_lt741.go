/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:57 PM  4/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

func cherryPickup(grid [][]int) int {
	n := len(grid)
	tiles := make([][][]*int, n)

	for i := range tiles {
		tiles[i] = make([][]*int, n)
		for j := range tiles[i] {
			tiles[i][j] = make([]*int, n)
			for k := range tiles[i][j] {
				tiles[i][j][k] = -1
			}
		}
	}
}

func dfsWithMemory(grid [][]int, x1, y1, x2 int) int {
	y2 := x1 + y1 - x2

}