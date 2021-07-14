/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  1:48 PM  14/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"math"
	"sort"
)

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	data := append(sort.IntSlice{}, nums1...)
	data.Sort()
	sum := 0
	m := 0
	MOD := int(1e9 + 7)
	for i := 0; i < len(nums1); i++ {
		diff := int(math.Abs(float64(nums1[i] - nums2[i])))
		sum = (sum + diff) % MOD
		j := data.Search(nums2[i])
		if j < len(nums1) {
			m = int(math.Max(float64(m), float64(diff - (data[j] - nums2[i]))))
		}
		if j > 0 {
			m = int(math.Max(float64(m), float64(diff - (nums2[i] - data[j - 1]))))
		}
	}
	return (sum - m + MOD) % MOD
}
