/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: tree_vertical_traverse_lt987.g0
 * @Version: 1.0.0
 * @Data: 2021/7/31 5:01 PM
 */
package leetcode

import (
	"math"
	"sort"
)

type verticalNode struct {
	col int
	row int
	val int
}

func verticalTraversal(root *TreeNode) [][]int {
	data := []*verticalNode{}

	var dfsTraverse func(node *TreeNode, row, col int)
	dfsTraverse = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		data = append(data, &verticalNode{col, row, node.Val})
		dfsTraverse(node.Left, row+1, col-1)
		dfsTraverse(node.Right, row+1, col+1)
	}
	dfsTraverse(root, 0, 0)

	sort.Slice(data, func(i, j int) bool {
		return data[i].col < data[j].col || (data[i].col == data[j].col && data[i].row < data[j].row) ||
			(data[i].col == data[j].col && data[i].row == data[j].row && data[i].val < data[j].val)
	})

	var res [][]int
	lastCol := math.MinInt32
	for i := range data {
		if lastCol != data[i].col {
			res = append(res, []int{})
			lastCol = data[i].col
		}
		res[len(res) - 1] = append(res[len(res) - 1], data[i].val)
	}
	return res
}
