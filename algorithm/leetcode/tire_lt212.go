/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: tire_lt212.go
 * @Version: 1.0.0
 * @Data: 2021/9/16 9:18 AM
 */
package leetcode

type tire struct {
	children [26]*tire
	word string
}

func (t *tire) insertWord(word string) {
	node := t
	for _, c := range word {
		ch := c - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &tire{}
		}
		node = node.children[ch]
	}
	node.word = word
}

var dir = []struct{x, y int} {
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func findWords(board [][]byte, words []string) []string {
	root := &tire{}

	for _, s := range words {
		root.insertWord(s)
	}

	cache := make(map[string]bool)
	var dfs func(i, j int, root *tire)
	dfs = func(i, j int, root *tire) {
		ch := board[i][j]
		index := ch - 'a'
		node := root
		if node.children[index] == nil {
			return
		}

		node = node.children[index]
		if node.word != "" {
			cache[node.word] = true
		}

		board[i][j] = '.'
		for _, d := range dir {
			ii := i + d.x
			jj := j + d.y
			if ii >= 0 && ii < len(board) && jj >= 0 && jj < len(board[0]) && board[ii][jj] != '.' {
				dfs(ii, jj, node)
			}
		}
		board[i][j] = ch
	}

	for i := range board {
		for j := range board[i] {
			dfs(i, j, root)
		}
	}

	ans := make([]string, 0)
	for k, _ := range cache {
		ans = append(ans, k)
	}
	return ans
}
