/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  6:42 PM  13/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

func MergeSort(nums []int) []int {
	n := len(nums)

	if n == 1 {
		return nums
	}

	n /= 2
	l := MergeSort(nums[:n])
	r := MergeSort(nums[n:])
	return Merge(l, r)
}

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l) + len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}