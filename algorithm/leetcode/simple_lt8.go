/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:09 AM  23/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	flag := 0
	result := 0
	i := 0

	for i < len(s) && s[i] == ' ' {
		i++
		continue
	}
	if i < len(s) && s[i] == '-' {
		flag = 1
		i++
	} else if i < len(s) && s[i] == '+' {
		i++
	}
	for i < len(s) {
		v := rune(s[i])
		if unicode.IsDigit(v) {
			r := int(v) - '0'
			if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && r > 7) {
				if flag == 0 {
					return math.MaxInt32
				}
				return math.MinInt32
			}
			result = result*10 + r
			i++
		} else {
			break
		}
	}

	if flag == 0 {
		return result
	}
	return -result
}
