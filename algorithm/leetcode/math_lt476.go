/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: math_lt476.go
 * @Version: 1.0.0
 * @Data: 2021/10/18 2:32 PM
 */
package leetcode

func findComplement(num int) int {
	var t int
	for i := 0; i < 32; i++ {
		if num < (1 << i) {
			t = i
			break
		}
	}
	return num ^ ((1 << t) - 1)
}
