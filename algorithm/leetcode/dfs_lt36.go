/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dfs_lt36.go
 * @Version: 1.0.0
 * @Data: 2021/9/17 10:18 AM
 */
package leetcode

func isValidSudoku(board [][]byte) bool {
	m := len(board)
	n := len(board[0])
	cache := make(map[int]map[int]bool)
	rowCache := make(map[int]map[int]bool)
	columnCache := make(map[int]map[int]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}
			if !checkRow(i, j, rowCache, board) {
				return false
			}
			if !checkColumn(i, j, columnCache, board) {
				return false
			}
			if !checkGroup(i, j, cache, board) {
				return false
			}
		}
	}
	return true
}

// calc row, check valid
func checkRow(i, j int, rowCache map[int]map[int]bool, board [][]byte) bool {
	v := int(board[i][j] - '0')
	if vv, ok := rowCache[i]; ok {
		if b, ok := vv[v]; ok && b {
			return false
		}
	} else {
		rowCache[i] = make(map[int]bool)
	}
	rowCache[i][v] = true
	return true
}

// calc column, check valid
func checkColumn(i, j int, columnCache map[int]map[int]bool, board [][]byte) bool {
	v := int(board[i][j] - '0')
	if vv, ok := columnCache[j]; ok {
		if _, ok := vv[v]; ok {
			return false
		}
	} else {
		columnCache[j] = make(map[int]bool)
	}
	columnCache[j][v] = true
	return true
}

// calc group, check valid
func checkGroup(i, j int, cache map[int]map[int]bool, board [][]byte) bool {
	index := calcGroup(i, j)
	v := int(board[i][j] - '0')
	if vv, ok := cache[index]; ok {
		if _, ok := vv[v]; ok {
			return false
		}
	} else {
		cache[index] = make(map[int]bool)
	}
	cache[index][v] = true
	return true
}

func calcGroup(i, j int) int {
	return (i / 3) * 3 + j / 3
}
