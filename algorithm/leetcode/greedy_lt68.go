/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: greedy_lt68.go
 * @Version: 1.0.0
 * @Data: 2021/9/9 3:48 PM
 */
package leetcode

import "strings"

/*
我们分以下三种情况讨论：

当前行是最后一行：单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格；
当前行不是最后一行，且只有一个单词：该单词左对齐，在行末填充空格；
当前行不是最后一行，且不只一个单词：设当前行单词数为  numWords，空格数为 numSpaces，我们需要将空格均匀分配在单词之间，则单词之间应至少有
	avgSpace = floor(numSpace/(numWords-1))
    extraSpace = numSpace mod (numWords-1)
 */

func fullJustify(words []string, maxWidth int) []string {
	index := 0
	n := len(words)
	ans := make([]string, 0)
	for {
		left := index
		size := 0
		for index < n && size + len(words[index]) + index - left <= maxWidth {
			size += len(words[index])
			index++
		}

		if index == n {
			//lastLine()
			s := strings.Join(words[left:], " ")
			s = s + blank(maxWidth - len(s))
			ans = append(ans, s)
			break
		}

		numberWords := index - left
		numberSpace := maxWidth - size

		if numberWords == 1 {
			//OneWord()
			s := words[left]
			s = s + blank(numberSpace)
			ans = append(ans, s)
		} else {
			//MultiWord()
			avgSpace := numberSpace / (numberWords - 1)
			extraSpace := numberSpace % (numberWords - 1)
			s1 := strings.Join(words[left:left + extraSpace + 1], blank(avgSpace + 1))
			s2 := strings.Join(words[left+extraSpace+1:index], blank(avgSpace))
			ans = append(ans, s1 + blank(avgSpace) + s2)
		}
	}
	return ans
}

func blank(n int) string {
	return strings.Repeat(" ", n)
}