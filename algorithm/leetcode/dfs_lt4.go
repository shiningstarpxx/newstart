/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:43 AM  3/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	f := findK(nums1, nums2, 0, len(nums1) - 1, 0, len(nums2) - 1, (len(nums1) + len(nums2) + 1)/2)
	s := findK(nums1, nums2, 0, len(nums1) - 1, 0, len(nums2) - 1, (len(nums1) + len(nums2) + 2)/2)
	fmt.Printf("%d, %d\n", f, s)
	return float64(f + s) / 2
}

func findK(num1, num2 []int, left1, right1, left2, right2, k int) int {
	len1 := right1 - left1 + 1
	len2 := right2 - left2 + 1
	if len1 > len2 {
		return findK(num2, num1, left2, right2, left1, right1, k)
	}

	// edge case, direct return
	if len1 == 0 {
		return num2[left2 + k - 1]
	}
	if k == 1 {
		return int(math.Min(float64(num1[left1]), float64(num2[left2])))
	}

	// downsize 2/k
	k1 := left1 + int(math.Min(float64(len1), float64(k / 2 ))) - 1
	k2 := left2 + int(math.Min(float64(len2), float64(k / 2 ))) - 1
	if num1[k1] <= num2[k2] {
		return findK(num1, num2, k1 + 1, right1, left2, right2, k - (k1 - left1 + 1))
	} else {
		return findK(num1, num2, left1, right1, k2 + 1, right2, k - (k2 - left2 + 1))
	}
}