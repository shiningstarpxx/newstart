/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:29 AM  6/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

// TODO(michael): In C++, upper_bound or lower_bound is generic function for all kinds type,
// also support self-defined type. We Would like to implement suck kinds of algorithms for golang

/*
UppperBound returns an iterator pointing to the first element in the range [first,last)
which has a value greater than ‘val’.
*/
func UppperBound(arr []int, target int) int {
	l, r := 0, len(arr)
	for ; l < r; {
		mid := l + (r-l)/2
		if arr[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

/*
LowerBound returns an iterator pointing to the first element in the range [first,last)
which has a value not less than ‘val’.
*/
func LowerBound(arr []int, target int) int {
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
