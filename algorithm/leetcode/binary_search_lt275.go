/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  12:15 PM  12/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

// h 指数的定义: “h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）
// 总共有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）"

func hIndex(citations []int) int {
	left := 0
	right := len(citations) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if citations[mid] >= len(citations) - mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return len(citations) - left
}
