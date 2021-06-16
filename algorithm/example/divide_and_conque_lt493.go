/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  4:36 PM  16/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package example

func reversePairsDC(nums []int) int {
	return reversePairsHelper(nums, 0, len(nums))
}

func reversePairsHelper(nums []int, l, r int) int {
	if l == r {
		return 0
	}

	mid := l + (r - l) / 2
	ret := reversePairsHelper(nums, l, mid) + reversePairsHelper(nums, mid + 1, r)

	// after divide, the left and right are sorted
	for i := l; i <= mid; i++ {
		arr := nums[mid + 1: r + 1]
		index := UppperBound(arr, nums[i] / 2)
		ret += index - (r - mid)
	}

	return ret
}

func UppperBound(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func Merge2(l, r []int) []int {
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