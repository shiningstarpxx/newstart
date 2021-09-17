/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dfs_lt36_test.go
 * @Version: 1.0.0
 * @Data: 2021/9/17 10:48 AM
 */
package leetcode

import "testing"

func Test_calcGroup(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"9 * 9, first group",
			args{
				1, 2,
			},
			0,
		},
		{
			"9 * 9, fourth group",
			args{
				5, 5,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcGroup(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("calcGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"simple test",
			args{
				[][]byte{
					{'5','3','.','.','7','.','.','.','.'},
						{'6','.','.','1','9','5','.','.','.'},
						{'.','9','8','.','.','.','.','6','.'},
						{'8','.','.','.','6','.','.','.','3'},
						{'4','.','.','8','.','3','.','.','1'},
						{'7','.','.','.','2','.','.','.','6'},
						{'.','6','.','.','.','.','2','8','.'},
						{'.','.','.','4','1','9','.','.','5'},
						{'.','.','.','.','8','.','.','7','9'},

				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}