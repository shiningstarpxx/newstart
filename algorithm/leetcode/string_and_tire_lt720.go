/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:13 PM  4/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func longestWord(words []string) string {
	cache := map[string]int{}
	for _, v := range words {
		cache[v] = 1
	}

	r := ""
	for _, v := range words {
		if len(v) < len(r) || (len(v) == len(r) && v > r) {
			continue
		}
		s := ""
		i := 0
		for ; i < len(v); i++ {
			s += string(v[i])
			if _, ok := cache[s]; !ok {
				break
			}
		}
		if i == len(v) {
			r = v
		}
	}
	return r
}