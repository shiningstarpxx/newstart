/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: string_lt171.go
 * @Version: 1.0.0
 * @Data: 2021/7/30 5:08 PM
 */
package leetcode

func titleToNumber(columnTitle string) int {
	ans := 0
	for i := range columnTitle {
		ans = ans * 26 + int(columnTitle[i] - 'A') + 1
	}
	return ans
}