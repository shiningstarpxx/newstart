/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: bfs_lt847_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/6 9:20 AM
 */
package leetcode

import "testing"

func Test_shortestPathLength(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"simple test",
			args{
				[][]int{
					{1, 2, 3},
					{0},
					{0},
					{0},
				},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathLength(tt.args.graph); got != tt.want {
				t.Errorf("shortestPathLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
