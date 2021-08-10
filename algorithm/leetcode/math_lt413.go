/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: math_lt413.g
 * @Version: 1.0.0
 * @Data: 2021/8/10 4:49 PM
 */
package leetcode

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	ans := 0
	t := 0
	d := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i] - nums[i - 1] == d {
			t++
		} else {
			d = nums[i] - nums[i - 1]
			t = 0
		}
		ans += t
	}
	return ans
}
