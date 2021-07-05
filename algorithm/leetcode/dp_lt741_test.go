/*
* @Author Pei, Xingxin
* @Description //TODO $
* @Date $ $
* @Param $
 */
package leetcode

import "testing"

func Test_cherryPickup(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases
		{
			"simple test",
			args{
				[][]int{
					{0, 1, -1},
					{1, 0, -1},
					{1, 1, 1},
				},
			},
			5,
		}, // [[1,1,-1],[1,-1,1],[-1,1,1]]
		{
			"negative test",
			args{
				[][]int{
					{1, 1, -1},
					{1, -1, 1},
					{-1, 1, 1},
				},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cherryPickup(tt.args.grid); got != tt.want {
				t.Errorf("cherryPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}