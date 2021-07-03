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

func greedyCheck(s string) bool {
	// l, r表示「可能多余的左括号」，一个下界，一个上界，很直观。执行起来就是
	l, r := 0, 0
	for i := range s {
		if s[i] == '(' {
			l++
			r++
		} else if s[i] == ')' {
			if l > 0 {
				l--
			}
			r--
		} else {
			if l > 0 {
				l--
			}
			r++
		}
		if r < 0{
			return false
		}
	}
	return l == 0
}