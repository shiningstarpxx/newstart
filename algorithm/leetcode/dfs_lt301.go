/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:31 AM  1/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func removeInvalidParentheses(s string) []string {
	r := dfs(s, 0, 0, 0, "")
	return remove(r)
}

func dfs(s string, index, left, right int, r string) []string {
	if index == len(s) {
		if left == right {
			return []string{r}
		} else {
			return []string{}
		}
	}
	var res []string
	if s[index] == ')' {
		res = append(res, dfs(s, index + 1, left, right, r)...)
		if right < left {
			res = append(res, dfs(s, index + 1, left, right + 1, string(append([]byte(r), s[index])))...)
		}
	} else if s[index] == '(' {
		res = append(res, dfs(s, index + 1, left + 1, right, string(append([]byte(r), s[index])))...)
	} else {
		res = append(res, dfs(s, index + 1, left, right, string(append([]byte(r), s[index])))...)
	}
	return res
}

func remove(r []string) []string {
	var res []string
	m := make(map[string]int, len(r))
	for _, v := range r {
		if _, ok := m[v]; !ok {
			res = append(res, v)
			m[v] = 1
		}
	}
	if len(res) == 0 {
		res = append(res, "")
	}
	return res
}