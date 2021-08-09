/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: two_pointers_slow_fast_lt147.go
 * @Version: 1.0.0
 * @Data: 2021/8/7 7:16 PM
 */
package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
	    return false
    }
	slow := head
	fast := head.Next
	for slow != fast {
	    if fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return true
}
