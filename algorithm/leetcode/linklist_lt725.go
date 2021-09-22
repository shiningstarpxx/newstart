/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: linklist_lt725.go
 * @Version: 1.0.0
 * @Data: 2021/9/22 7:35 AM
 */
package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	len := calcLength(head)
	addition := len % k
	average := len / k
	index := 0
	res := []*ListNode{}
	t := head
	for index < k {
		tt := t
		var pre *ListNode
		c := 0
		for c < average {
			if t != nil {
				pre = t
				t = t.Next
			}
			c++
		}
		if addition > 0 {
			if t != nil {
				pre = t
				t = t.Next
			}
			addition--
		}
		if pre != nil {
			pre.Next = nil
		}
		res = append(res, tt)
		index++
	}
	return res
}

func calcLength(head *ListNode) int {
	len := 0
	t := head
	for t != nil {
		len++
		t = t.Next
	}
	return len
}
