/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: hashtable_lt447.go
 * @Version: 1.0.0
 * @Data: 2021/9/13 11:34 AM
 */
package leetcode

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for i := range points {
		cache := make(map[int]int)
		for j := range points {
			if i != j {
				d := (points[i][0] - points[j][0]) * (points[i][0] - points[j][0]) +
					(points[i][1] - points[j][1]) * (points[i][1] - points[j][1])
				cache[d]++
			}
		}
		for _, v := range cache {
			ans += v * (v - 1)
		}
	}
	return ans
}
