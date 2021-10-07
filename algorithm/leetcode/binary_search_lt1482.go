/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: binary_search_lt1482.go
 * @Version: 1.0.0
 * @Data: 2021/10/5 10:30 PM
 */
package leetcode

func minDays(bloomDay []int, m int, k int) int {
	l := len(bloomDay)
	if l < m * k {
		return -1
	}

	maxDay := 0
	for i := range bloomDay {
		maxDay = max1482(maxDay, bloomDay[i])
	}

	l, r := 1, maxDay
	ans := maxDay
	for l <= r {
		mid := l + (r - l) / 2
		if canMake(bloomDay, mid, m, k) {
			ans = min1482(ans, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}

func min1482(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1482(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canMake(bloomDay []int, mid int, m, k int) bool {
	count := 0
	banquetCount := 0
	for i := range bloomDay {
		if bloomDay[i] <= mid {
			banquetCount++
			if banquetCount == k {
				count++
				banquetCount = 0
			}
		} else {
			banquetCount = 0
		}
	}
	return count >= m
}