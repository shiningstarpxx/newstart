/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_decode_way_lt639.go
 * @Version: 1.0.0
 * @Data: 2021/9/27 8:58 AM
 */
package leetcode

const (
	mod int = 1e9 + 7
)

func numDecodings(s string) int {
	n := len(s)

	pp, p := 0, 1

	calcOne := func (b byte) int {
		if b == '0' {
			return 0
		}
		if b == '*' {
			return 9 * p
		}
		return p
	}

	// ** *[] []* [][]
	calcTwo639 := func (a, b byte) int {
		if a == '*' && b == '*' {
			return 15 * pp
		}

		if a == '*' {
			if b >= '0' && b <= '6' {
				return 2 * pp
			}
			return pp
		}

		if b == '*' {
			if a == '1' {
				return 9 * pp
			}
			if a == '2' {
				return 6 * pp
			}
			return 0
		}

		if a != '0' && (a - '0') * 10 + (b - '0') <= 26 {
			return pp
		}
		return 0
	}

	for i := 0; i < n; i++ {
		c := calcOne(s[i]) % mod
		if i > 0 {
			c = (c + calcTwo639(s[i-1], s[i])) % mod
		}
		pp, p = p, c
	}
	return p
}
