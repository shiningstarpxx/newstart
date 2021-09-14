/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: string_lt524.go
 * @Version: 1.0.0
 * @Data: 2021/9/14 9:06 AM
 */
package leetcode

import "sort"

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j]) ||
			(len(dictionary[i]) == len(dictionary[j]) && dictionary[i] < dictionary[j])
	})

	for _, v := range dictionary {
		i, j := 0, 0
		for i < len(s) {
			if s[i] == v[j] {
				i++
				j++
			} else {
				i++
			}
			if j == len(v) {
				return v
			}
		}
	}
	return ""
}
