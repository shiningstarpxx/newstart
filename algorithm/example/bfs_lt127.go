/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  1:12 PM  26/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package example

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}
	cache := make(map[string]int, len(wordList))
	for _, v := range wordList {
		cache[v] = 0  // in fact, i is not used
	}

	var q []string
	q = append(q, beginWord)
	res := 0
	for len(q) > 0 {
		qq := q
		q = nil
		for len(qq) > 0 {
			s := qq[0]
			if s == endWord {
				return res + 1
			}
			qq = qq[1:]

			for i, v := range s {
				for c := 'a'; c <= 'z'; c++ {
					if v != c {
						ts := s[0:i] + string(c) + s[i+1:]
						if count, ok := cache[ts]; ok {
							if count == 0 {
								cache[ts] = 1
								q = append(q, ts)
							}
						}
					}
				}
			}
		}
		res++
	}
	return 0
}
