/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  2:30 PM  17/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"sort"
	"strconv"
	"unicode"
)

func countOfAtoms(formula string) string {
	c, _ := countOfAtomsDFS(formula, 0)
	type pair struct {
		atom string
		num  int
	}
	pairs := make([]pair, 0, len(c))
	for k, v := range c {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].atom < pairs[j].atom })

	ans := []byte{}
	for _, p := range pairs {
		ans = append(ans, p.atom...)
		if p.num > 1 {
			ans = append(ans, strconv.Itoa(p.num)...)
		}
	}
	return string(ans)
}

func countOfAtomsDFS(formula string, index int) (map[string]int, int) {
	c := map[string]int{}

	for index < len(formula) {
		if formula[index] == '(' {
			index++
			cc, index1 := countOfAtomsDFS(formula, index)
			count, index2 := getCount(formula, index1)
			for k, v := range cc {
				c[k] += v * count
			}
			index = index2
		} else if formula[index] == ')' {
			index++
			return c, index
		} else {
			name, index1 := getName(formula, index)
			count, index2 := getCount(formula, index1)
			c[name] += count
			index = index2
		}
	}
	return c, index
}

func getName(formula string, index int) (string, int)  {
	start := index
	index++ // 扫描，跳过首字母
	for index < len(formula) && unicode.IsLower(rune(formula[index])) {
		index++ // 扫描首字母后的小写字母
	}
	return formula[start:index], index
}

func getCount(formula string, index int) (int, int)  {
	num := 0
	if index == len(formula) || !unicode.IsDigit(rune(formula[index])) {
		return 1, index // 不是数字，视作 1
	}
	for ; index < len(formula) && unicode.IsDigit(rune(formula[index])); index++ {
		num = num*10 + int(formula[index]-'0') // 扫描数字
	}
	return num, index
}