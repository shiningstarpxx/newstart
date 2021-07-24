/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: brute_force_lt1736.go
 * @Version: 1.0.0
 * @Data: 2021/7/24 11:55 AM
 */
package leetcode

func maximumTime(time string) string {
	// 分钟比较无脑，第一个永远是'5'，第二个永远是'9'
	s := ""
	if time[0] == '?' {
		if time[1] == '?' || time[1] <= '3' {
			s += "2"
		} else {
			s += "1"
		}
	} else {
		s += string(time[0])
	}

	if time[1] == '?' {
		if s[0] == '2' {
			s += "3"
		} else {
			s += "9"
		}
	} else {
		s += string(time[1])
	}

	s += ":"

	if time[3] == '?' {
		s += "5"
	} else {
		s += string(time[3])
	}

	if time[4] == '?' {
		s += "9"
	} else {
		s += string(time[4])
	}
	return s
}
