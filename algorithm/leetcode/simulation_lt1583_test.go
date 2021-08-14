/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: simulation_lt1583_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/14 5:14 PM
 */
package leetcode

import "testing"

func Test_unhappyFriends(t *testing.T) {
	type args struct {
		n           int
		preferences [][]int
		pairs       [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple test",
			args{
				4,
				[][]int{
					{1, 2, 3},
					{3, 2, 0},
					{3, 1, 0},
					{1, 2, 0},
				},
				[][]int{
					{0, 1},
					{2, 3},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unhappyFriends(tt.args.n, tt.args.preferences, tt.args.pairs); got != tt.want {
				t.Errorf("unhappyFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}
