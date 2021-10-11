/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: simulation_lt273.go
 * @Version: 1.0.0
 * @Data: 2021/10/11 9:45 PM
 */
package leetcode

import "strings"

var (
	singles   = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens      = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands = []string{"", "Thousand", "Million", "Billion"}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	sb := &strings.Builder{}
	var cacl func(int)
	cacl = func(i int) {
		switch {
		case i == 0:
		case i < 10:
			sb.WriteString(singles[i])
			sb.WriteByte(' ')
		case i < 20:
			sb.WriteString(teens[i - 10])
			sb.WriteByte(' ')
		case i < 100:
			sb.WriteString(tens[i / 10])
			sb.WriteByte(' ')
			cacl(i % 10)
		default:
			sb.WriteString(singles[i / 100])
			sb.WriteString(" Hundred ")
			cacl(i % 100)
		}
	}

	for i, unit := 3, int(1e9); i >= 0; i-- {
		if curNum := num / unit; curNum > 0 {
			num -= curNum * unit
			cacl(curNum)
			sb.WriteString(thousands[i])
			sb.WriteByte(' ')
		}
		unit /= 1000
	}

	return strings.TrimSpace(sb.String())
}