/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: sort_squeeze_lt881.go
 * @Version: 1.0.0
 * @Data: 2021/8/26 9:44 AM
 */
package leetcode

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	ans := 0
	l, r := 0, len(people) - 1
	for l <= r {
		if people[l] + people[r] > limit {
			r--
		} else {
			l++
			r--
		}
		ans++
	}
	return ans
}
