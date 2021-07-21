/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: offer_52.go
 * @Version: 1.0.0
 * @Data: 2021/7/21 11:53 AM
 */
package interview

type ListNode struct {
	Val int
	Next *ListNode
}

// assume
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	return isIntersect(headA, headB, headA, headB, true, true)
}

func isIntersect(headA, headB, oa, ob *ListNode, swapa, swapb bool) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	if headA == headB {
		return headA
	}
	if headA.Next != nil && headB.Next != nil {
		return isIntersect(headA.Next, headB.Next, oa, ob, swapa, swapb)
	}
	nexta, nextb := headA.Next, headB.Next
	if nexta == nil && swapa {
		swapa = false
		nexta = ob
	}
	if nextb == nil && swapb {
		swapb = false
		nextb = oa
	}
	if nexta != nil && nextb != nil {
		return isIntersect(nexta, nextb, oa, ob, swapa, swapb)
	}
	return nil
}

func getIntersectionNodeOrigin(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}