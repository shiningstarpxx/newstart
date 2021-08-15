/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: simulation_lt1583.go
 * @Version: 1.0.0
 * @Data: 2021/8/14 4:48 PM
 */
package leetcode

// pair[x][y] , pair[u][v]
// x 与 u 的亲近程度胜过 x 与 y，   (index[x][u] < index[x][y])
// u 与 x 的亲近程度胜过 u 与 v     (index[u][x] < index[u][v]
func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	index := make([][]int, n)
	for i := range index {
		index[i] = make([]int, n)
	}

	// 记录每对的索引 -- hash table 化
	for i, v := range preferences {
		for j := range v {
			index[i][v[j]] = j
		}
	}

	// 记录对--hash table 化
	match := make([]int, n)
	for _, v := range pairs {
		match[v[0]] = v[1]
		match[v[1]] = v[0]
	}

	ans := 0
	for x, y := range match {
		i := index[x][y]  // index值, xy 配对了，找到这个索引值，然后比他小的值都有可能unhappy
		for _, u := range preferences[x][:i] {
			v := match[u]  // 找到比他小的index，判断反过来是不是也unhappy
			if index[u][v] > index[u][x] {
				ans++
				break
			}
		}
	}
	return ans
}
