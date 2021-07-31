/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: tree_lt863.go
 * @Version: 1.0.0
 * @Data: 2021/7/29 12:03 AM
 */
package leetcode

// todo(michael): 增加一个能力，从任意一个节点遍历整个树，只走一次
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	cache := map[int]*TreeNode{}
	ans := []int{}

	var findParent func(*TreeNode)
	findParent = func(root *TreeNode) {
		if root.Left != nil {
			cache[root.Left.Val] = root
			findParent(root.Left)
		}
		if root.Right != nil {
			cache[root.Right.Val] = root
			findParent(root.Right)
		}
	}
	findParent(root)

	var findAns func(root *TreeNode, from *TreeNode, depth int)
	findAns = func(root *TreeNode, from *TreeNode, depth int) {
		if root == nil {
			return
		}

		if depth == k {
			ans = append(ans, root.Val)
			return
		}

		if root.Left != from {
			findAns(root.Left, root, depth + 1)
		}
		if root.Right != from {
			findAns(root.Right, root, depth + 1)
		}
		if cache[root.Val] != from {
			findAns(cache[root.Val], root, depth + 1)
		}
	}
	findAns(target, nil, 0)
	return ans
}
