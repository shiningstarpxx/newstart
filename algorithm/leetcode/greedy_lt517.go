/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: greedy_lt517.go
 * @Version: 1.0.0
 * @Data: 2021/9/29 10:11 PM
 */
package leetcode

func findMinMoves(machines []int) int {
	sum := 0
	for i := range machines {
		sum += machines[i]
	}

	n := len(machines)
	if sum % n != 0 {
		return -1
	}

	average := sum / n
	sum = 0
	num := 0
	ans := 0
	for i := range machines {
		num = machines[i] - average
		sum += num
		ans = max517(ans, max517(abs517(sum), num))
	}
	return ans
}

func max517(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs517(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
