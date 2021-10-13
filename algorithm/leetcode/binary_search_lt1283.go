/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: binary_search_lt1283.go
 * @Version: 1.0.0
 * @Data: 2021/10/13 9:19 PM
 */
package leetcode

func smallestDivisor(nums []int, threshold int) int {
	l, r := 1, getMax(nums)
	ans := 0

	for l <= r {
		mid := l + (r - l) >> 1
		sum := 0
		for i := 0; i < len(nums); i++ {
			sum += (nums[i] - 1) / mid + 1
		}
		if sum <= threshold {
			ans = mid
			r = mid -1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func getMax(num []int) int {
	res := 0
	for i := 0; i < len(num); i++ {
		res = max1283(res, num[i])
	}
	return res
}

func max1283(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
