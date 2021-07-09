/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  1:22 PM  9/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func numSubarraysWithSum(nums []int, goal int) int {
	sum := 0
	cache := map[int]int{}
	count := 0
	for _, v := range nums {
		cache[sum]++
		sum += v
		count += cache[sum - goal]
	}
	return count
}