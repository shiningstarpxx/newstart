/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:58 AM  3/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func checkValidString(s string) bool {
	return stackCheck(s)
}

func stackCheck(s string) bool {
	var left []int
	var right []int

	for i := range s {
		if s[i] == '(' {
			left = append(left, i)
		} else if s[i] == '*'{
			right = append(right, i)
		} else {
			if len(left) > 0 {
				left = left[:len(left) - 1]
			} else if len(right) > 0 {
				right = right[:len(right) - 1]
			} else {
				return false
			}
		}
	}
	for len(left) > 0 && len(right) > 0 {
		if left[len(left) - 1] > right[len(right) - 1] {
			return false
		}
		left = left[:len(left) - 1]
		right = right[:len(right) - 1]
	}
	return len(left) == 0
}
