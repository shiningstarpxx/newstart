/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: hash_map_lt1436.go
 * @Version: 1.0.0
 * @Data: 2021/10/1 10:10 AM
 */
package leetcode

func destCity(paths [][]string) string {
	from := make(map[string]int)
	to := make(map[string]int)
	for i := range paths {
		from[paths[i][0]] = 1
		to[paths[i][1]] = 1
	}

	for k, _ := range from {
		to[k] = 0
	}

	for k, v := range to {
		if v == 1 {
			return k
		}
	}
	return ""
}

func destCityHelper(paths [][]string) string {
	cache := make(map[string]bool)
	for i := range paths {
		cache[paths[i][0]] = true
	}
	for i := range paths {
		if !cache[paths[i][1]] {
			return paths[i][1]
		}
	}
	return ""
}