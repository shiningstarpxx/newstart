/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  4:50 PM  16/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package interview

func search(nums []int, target int) int {
	return uppperBound(nums, target) - lowerBound(nums, target)
}

func uppperBound(arr []int, target int) int {
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

func lowerBound(arr []int, target int) int {
	l, r := 0, len(arr)
	for ; l < r; {
		mid := l + (r-l)/2
		if arr[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
