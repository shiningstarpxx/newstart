/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: strings_lt165.go
 * @Version: 1.0.0
 * @Data: 2021/9/1 10:31 AM
 */
package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i, j := 0, 0; i < len(v1) || j < len(v2); i, j = i+1, j+1 {
		x, y := 0, 0
		if i < len(v1) {
			x, _ = strconv.Atoi(v1[i])
		}
		if j < len(v2) {
			y, _ = strconv.Atoi(v2[j])
		}

		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
	}
	return 0
}
