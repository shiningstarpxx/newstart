/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  6:30 PM  27/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package example

func calcOne91(c byte) int {
	if c == '0' {
		return 0
	}
	return 1
}

func calcTwo91(b, c byte) int {
	if b == '0' {
		return 0
	}
	if b == '1' {
		return 1
	}
	if b == '2' {
		if c <= '6' {
			return 1
		}
		return 0
	}
	return 0
}

func numDecodings91(s string) int {
	if len(s) == 0 {
		return 0
	}
	first, second := 1, calcOne91(s[0])
	for i := 1; i < len(s); i++ {
		first, second = second, first * calcTwo91(s[i-1], s[i]) + second * calcOne91(s[i])
	}
	return second
}
