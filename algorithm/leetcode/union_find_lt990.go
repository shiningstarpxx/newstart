/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  5:35 PM  30/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

/*
给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。 

示例 1：

输入：["a==b","b!=a"]
输出：false
解释：如果我们指定，a = 1 且 b = 1，那么可以满足第一个方程，但无法满足第二个方程。没有办法分配变量同时满足这两个方程。
示例 2：

输入：["b==a","a==b"]
输出：true
解释：我们可以指定 a = 1 且 b = 1 以满足满足这两个方程。
示例 3：

输入：["a==b","b==c","a==c"]
输出：true
示例 4：

输入：["a==b","b!=c","c==a"]
输出：false
示例 5：

输入：["c==c","b==d","x!=z"]
输出：true

提示：
1 <= equations.length <= 500
equations[i].length == 4
equations[i][0] 和 equations[i][3] 是小写字母
equations[i][1] 要么是 '='，要么是 '!'
equations[i][2] 是 '='
 */

import (
	"golang_source_analyses/algorithm/slice"
	"strings"
)

// 跟下面的版本比，两次循环，但是代码简洁程度更高
func equationsPossible(equations []string) bool {
	uf := algorithm.NewUnionFind(26)
	for _, v := range equations {
		if strings.Contains(v, "==") {
			r := strings.Split(v, "==")
			uf.Union(int(r[0][0]-'a'), int(r[1][0]-'a'))
		}
	}
	for _, v := range equations {
		if strings.Contains(v, "!=") {
			r := strings.Split(v, "!=")
			if uf.IsConnected(int(r[0][0]-'a'), int(r[1][0]-'a')) {
				return false
			}
		}
	}
	return true
}

func equationsPossibleV1(equations []string) bool {
	uf := algorithm.NewUnionFind(26)
	t := make([][]int, 0)
	for _, v := range equations {
		if strings.Contains(v, "==") {
			r := strings.Split(v, "==")
			uf.Union(int(r[0][0]-'a'), int(r[1][0]-'a'))
		} else {
			r := strings.Split(v, "!=")
			tt := make([]int, 2)
			tt[0] = int(r[0][0]-'a')
			tt[1] = int(r[1][0]-'a')
			t = append(t, tt)
		}
	}
	for _, v := range t {
		if uf.IsConnected(v[0], v[1]) {
			return false
		}
	}
	return true
}

/*
执行用时： 4 ms, 在所有 Go 提交中击败了75.00%的用户 (版本V1击败了 100%)
内存消耗：3 MB, 在所有 Go 提交中击败了19.83%的用户
 */