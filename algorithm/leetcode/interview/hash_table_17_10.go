/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  1:17 PM  9/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package interview

func majorityElement(nums []int) int {
	cache := map[int]int{}

	for _, v := range nums {
		cache[v]++
		if cache[v] > len(nums) / 2 {
			return v;
		}
	}

	return -1
}
