/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: math_lt29.go
 * @Version: 1.0.0
 * @Data: 2021/10/12 6:53 PM
 */
package leetcode

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 {
		if divisor == -1 {
			return math.MaxInt32
		}
		if divisor == 1 {
			return math.MinInt32
		}
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}

	flag := false
	if dividend > 0 {
		dividend = -dividend
		flag = !flag
	}
	if divisor > 0 {
		divisor = -divisor
		flag = !flag
	}
	ans := 0

	t := divisor
	cache := make([]int, 0)
	for t >= dividend {
		cache = append(cache, t)
		t = t << 1;
	}

	for i := len(cache) - 1; i >= 0; i-- {
		if cache[i] >= dividend {
			ans += 1 << i
			dividend -= cache[i]
		}
	}

	if flag {
		return -ans;
	}
	return ans;
}
