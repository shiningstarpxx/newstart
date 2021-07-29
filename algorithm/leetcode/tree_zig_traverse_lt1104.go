/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: tree_zig_traverse_lt1104.go
 * @Version: 1.0.0
 * @Data: 2021/7/29 2:05 PM
 */
package leetcode

func reverseLabel(row, label int) int {
	if row % 2 != 0 {
		return label
	}
	return 1 << (row - 1) + 1 << row - 1 - label
}

func calcRow(label int) int {
	row := 0
	for label > 0 {
		label /= 2
		row++
	}
	return row
}

func pathInZigZagTree(label int) (ans []int) {
	row := calcRow(label)

	label = reverseLabel(row, label)
	for row > 0 {
		ans = append(ans, reverseLabel(row, label))
		row--
		label /= 2
	}
	reverseSlice(ans)
	return ans
}
