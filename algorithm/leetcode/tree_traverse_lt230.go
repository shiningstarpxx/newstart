/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: tree_traverse_lt230.go
 * @Version: 1.0.0
 * @Data: 2021/10/18 11:12 AM
 */
package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack) - 1], stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return 0
}
