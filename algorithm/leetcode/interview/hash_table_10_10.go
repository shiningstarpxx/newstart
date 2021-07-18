/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  1:39 PM  18/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package interview

import "sort"

func groupAnagrams(strs []string) [][]string {
	return groupAnagramsWithSort(strs)
}

func groupAnagramsWithSort(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func groupAnagramsWithHashTable(strs []string) [][]string {
	c := map[[26]int][]string{}

	for _, v := range strs {
		cc := [26]int{}
		for i := range v {
			cc[v[i] - 'a']++
		}
		c[cc] = append(c[cc], v)
	}

	ans := make([][]string, 0, len(c))
	for _, v := range c {
		ans = append(ans, v)
	}
	return ans
}