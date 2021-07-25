/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: hash_table_lt1743.go
 * @Version: 1.0.0
 * @Data: 2021/7/25 4:08 PM
 */
package leetcode

func restoreArray(adjacentPairs [][]int) []int {
	// [1, 2], [2, 3] => [1, 2, 3], result size is adjacentPairs.size() + 1
	n := len(adjacentPairs) + 1
	ans := make([]int, n)
	cache := map[int][]int {}
	for _, e := range adjacentPairs {
		cache[e[0]] = append(cache[e[0]], e[1])
		cache[e[1]] = append(cache[e[1]], e[0])
	}

	for k, v := range cache {
		if len(v) == 1 {
			ans[0] = k
			ans[1] = v[0]
			break
		}
	}

	for i := 2; i < n; i++ {
		v := cache[ans[i - 1]]
		if v[0] == ans[i - 2] {
			ans[i] = v[1]
		} else {
			ans[i] = v[0]
		}
	}
	return ans
}
