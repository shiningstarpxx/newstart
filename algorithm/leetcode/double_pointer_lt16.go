/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: double_pointer_lt16.go
 * @Version: 1.0.0
 * @Data: 2021/7/24 12:34 PM
 */
package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := math.MaxInt32
	res := 0
	for i := 0; i < len(nums) - 2; i++ {
		m, n := i + 1, len(nums) - 1
		for m < n {
			diff := target - (nums[i] + nums[m] + nums[n])
			// fmt.Printf("%d, %d, %d\n", nums[i] + nums[m] + nums[n], ans, diff)
			if diff == 0 {
				return target
			}
			if ans > int(math.Abs(float64(diff))) {
				ans = int(math.Abs(float64(diff)))
				res = nums[i] + nums[m] + nums[n]
			}
			if diff > 0 {
				m++
			} else {
				n--
			}
		}
	}
	return res
}
