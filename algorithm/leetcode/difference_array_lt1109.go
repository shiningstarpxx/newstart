/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: difference_array_lt1109.go
 * @Version: 1.0.0
 * @Data: 2021/8/31 3:56 PM
 */
package leetcode

func corpFlightBookings(bookings [][]int, n int) []int {
	diffrence := make([]int, n)

	for _, v := range bookings {
		diffrence[v[0] - 1] += v[2]
		if v[1]< n {
			diffrence[v[1]] -= v[2]
		}
	}

	for i := 1; i < n; i++ {
		diffrence[i] += diffrence[i-1]
	}
	return diffrence
}
