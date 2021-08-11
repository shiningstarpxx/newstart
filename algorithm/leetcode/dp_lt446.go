/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_lt446.go
 * @Version: 1.0.0
 * @Data: 2021/8/11 11:42 PM
 */
package leetcode

func numberOfArithmeticSlices2(nums []int) int {
	cache := make([]map[int]int, len(nums))
	ans := 0
	for i, x := range nums {
		cache[i] = map[int]int{}
		for j, y := range nums[:i] {
			d := x - y
			cnt := cache[j][d]
			ans += cnt
			cache[i][d] += cnt + 1
		}
	}
	return ans
}
