/**
 * @author peixingxin
 * @description //TODO
 * @date 3:57 PM 2021/7/19
 * @file sliding_window_lt1838.go
 **/

package leetcode

import (
	"math"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	// 7, 6, 5, 4, 3, 2, 1
	sort.Ints(nums)
	// 1, 2, 3, 4, 5, 6, 7
	ans := 0
	for l, r, total := 0, 1, 0; r < len(nums); r++ {
		// 1, 2, 3, 4, 5, 6, 7
		//    l    r-1 r
		//         r-1 是上一个边界, 那个时候的，如果到了r，比之前是多了 (nums[r] - nums[r-1])
		total += (nums[r] - nums[r - 1]) * (r - l)
		for total > k {
			total -= nums[r] - nums[l]
			l++
		}
		ans = int(math.Max(float64(ans), float64(r - l + 1)))
	}
	return ans
}
