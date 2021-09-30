/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: math_lt223.go
 * @Version: 1.0.0
 * @Data: 2021/9/30 5:20 PM
 */
package leetcode

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	width := min223(bx2, ax2) - max223(bx1, ax1)
	height := min223(by2, ay2) - max223(by1, ay1)
	d := width * height
	if width < 0 || height < 0 {
		d = 0
	}
	return (ax2 - ax1) * (ay2 - ay1) + (bx2 - bx1) * (by2 - by1) - d
}

func min223(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max223(a, b int) int {
	if a > b {
		return a
	}
	return b
}
