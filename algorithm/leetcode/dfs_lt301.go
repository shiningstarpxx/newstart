/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:31 AM  1/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func removeInvalidParentheses(s string) []string {
	l, r := countInvalid(s)
	res := dfs(s, 0, l, r, "")
	return res
}

func countInvalid(s string) (int, int) {
	l, r := 0, 0
	for i := range s {
		if s[i] == '(' {
			l++
		} else if s[i] == ')' {
			if l > 0 {
				l--
			} else {
				r++
			}
		}
	}
	return l, r
}

func dfs(s string, index, left, right int, r string) []string {
	if index == len(s) {
		if left == 0 && right == 0 && valid(r) {
			return []string{r}
		}
		return []string{}
	}

	var res []string
	if s[index] == '(' {
		if left > 0 {
			res = append(res, dfs(s, index+1, left-1, right, r)...)
		}
	} else if s[index] == ')' {
		if right > 0 {
			res = append(res, dfs(s, index+1, left, right-1, r)...)
		}
	}
	res = append(res, dfs(s, index+1, left, right, string(append([]byte(r), s[index])))...)
	return res
}

func valid(s string) bool {
	l := 0
	for _, v := range s {
		if v == '(' {
			l++
		} else if v == ')' {
			l--
			if l < 0 {
				return false
			}
		}
	}
	return l == 0
}
