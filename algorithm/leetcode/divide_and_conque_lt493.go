/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  4:36 PM  16/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func reversePairsDC(nums []int) int {
	return reversePairsHelper(nums, 0, len(nums) - 1)
}

func reversePairsHelper(nums []int, l, r int) int {
	if l == r {
		return 0
	}

	mid := l + (r - l) / 2
	ret := reversePairsHelper(nums, l, mid) + reversePairsHelper(nums, mid + 1, r)

	i, ll, rr := 0, l, mid + 1
	for ll <= mid && rr <= r {
		if nums[ll] > 2 * nums[rr] {
			ret += mid - ll + 1
			rr++
		} else {
			ll++
		}
	}

	// merge
	sorted := make([]int, r - l + 1)
	i, ll, rr = 0, l, mid + 1
	for ll <= mid || rr <= r {
		if ll > mid {
			sorted[i] = nums[rr]
			rr++
		} else if rr > r {
			sorted[i] = nums[ll]
			ll++
		} else if nums[ll] <= nums[rr] {
			sorted[i] = nums[ll]
			ll++
		} else {
			sorted[i] = nums[rr]
			rr++
		}
		i++
	}
	for i := 0; i < len(sorted); i++ {
		nums[i + l] = sorted[i]
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