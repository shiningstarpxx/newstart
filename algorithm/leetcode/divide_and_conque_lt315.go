/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  3:21 PM  14/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func countSmaller(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	countSmallerDC(nums, 0, n - 1, res)
	return res
}

func countSmallerDC(nums []int, l, r int, res []int) []int {
	if l == r {
		return []int{nums[l]}
	}

	mid := l + (r - l) / 2
	left := countSmallerDC(nums, l, mid, res)
	right := countSmallerDC(nums, mid + 1, r, res)

	for l <= mid {
		rindex := LowerBound(right, nums[l])
		res[l] +=  rindex
		l++
	}
	return Merge(left, right)
}

func LowerBound(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
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
