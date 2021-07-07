/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  2:31 PM  7/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func countPairs(deliciousness []int) int {
	cache := map[int]int{}
	max := 0
	for _, v := range deliciousness {
		if max < v {
			max = v
		}
	}
	max <<= 1

	count := 0
	for _, v := range deliciousness {
		for sum := 1; sum <= max; sum <<= 1 {
			if _, ok := cache[sum - v]; ok {
				count = (count + cache[sum - v]) % (1e9 + 7)
			}
		}
		cache[v]++
	}
	return count
}
