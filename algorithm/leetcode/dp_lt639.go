/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:17 AM  27/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

// https://zxi.mytechroad.com/blog/dynamic-programming/leetcode-639-decode-ways-ii/

func calcOne(c byte) int {
	if c == '0' {
		return 0
	}
	if c == '*' {
		return 9
	}
	return 1
}

func calcTwo(b, c byte) int {
	if b == '*' {
		if c == '*' {  // 11 - 19, 21-26
			return 15
		}
		if c <= '6' {
			return 2
		}
		return 1
	}
	if b == '0' {
		return 0
	}
	if b == '1' {
		if c == '*' {
			return 9
		} else {
			return 1
		}
	}
	if b == '2' {
		if c == '*' {
			return 6
		} else if c <= '6' {
			return 1
		}
		return 0
	}
	return 0
}

func numDecodings639(s string) int {
	if len(s) == 0 {
		return 0
	}
	first, second := 1, calcOne(s[0])
	for i := 1; i < len(s); i++ {
		first, second = second, first * calcTwo(s[i - 1], s[i]) + second * calcOne(s[i])
		second %= 1e9 + 7
	}
	return second
}