/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: tree_lt671.go
 * @Version: 1.0.0
 * @Data: 2021/7/27 1:14 PM
 */
package leetcode

import "math"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	first, second, flag := math.MaxInt64, math.MaxInt64, true
	_, second, flag = traverse(root, first, second, flag)
	if flag {
		return -1
	}
	return second
}

func traverse(root *TreeNode, first, second int, flag bool) (int, int, bool) {
	if root == nil {
		return first, second, flag
	}

	if second > root.Val {
		if first > root.Val {
			second = first
			first = root.Val
		} else if first != root.Val {
			second = root.Val
			flag = false
		}
	}
	first, second, flag = traverse(root.Left, first, second, flag)
	first, second, flag = traverse(root.Right, first, second, flag)
	return first, second, flag
}
