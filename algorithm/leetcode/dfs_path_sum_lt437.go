/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dfs_path_sum_lt437.go
 * @Version: 1.0.0
 * @Data: 2021/9/28 3:56 PM
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
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := 0
	ans += dfsPath(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}

func dfsPath(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if targetSum - root.Val == 0 {
		res++
	}
	res += dfsPath(root.Left, targetSum-root.Val)
	res += dfsPath(root.Right, targetSum-root.Val)
	return res
}
