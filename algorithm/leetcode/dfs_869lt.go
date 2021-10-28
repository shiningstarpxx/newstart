/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dfs_869lt.go
 * @Version: 1.0.0
 * @Data: 2021/10/28 9:19 AM
 */
package leetcode

import "strconv"

var cache map[int]int

func reorderedPowerOf2(n int) bool {
	cache = make(map[int]int)
	strNum := strconv.Itoa(n)
	return dfs869(strNum, 0, 0)
}

func isPowerOfTwo(n int) bool {
	return n & (n - 1) == 0
}

func dfs869(strNum string, index, num int) bool {
	if index == len(strNum) {
		return isPowerOfTwo(num)
	}

	for i := 0; i < len(strNum); i++ {
		if (num == 0 && strNum[i] == '0') || cache[i] == 1 {
			continue
		}
		cache[i] = 1
		if dfs869(strNum, index + 1, num * 10 + int(strNum[i] - '0')) {
			return true
		}
		cache[i] = 0
	}
	return false
}
