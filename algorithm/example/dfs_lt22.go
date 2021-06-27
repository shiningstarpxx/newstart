/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:18 PM  25/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package example

func generateParenthesis(n int) []string {
	res := helper(n, n, "")
	return res
}

func helper(left, right int, s string) []string {
	var res []string
	if left == 0 && right == 0 {
		res = append(res, s)
		return res
	}
	if left > 0 {
		res = append(res, helper(left - 1, right, s + "(")...)
	}
	if right > left {
		res = append(res, helper(left, right - 1, s + ")")...)
	}
	return res
}