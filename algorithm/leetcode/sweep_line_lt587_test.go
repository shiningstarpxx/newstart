/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: sweep_line_lt587_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/3 2:16 PM
 */
package leetcode

import "testing"

func Test_findUnsortedSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple test",
			args{
				[]int{2,6,4,8,10,9,15},
			},
			5,
		},
		{
			"edge casee",
			args{
				[]int{1, 2, 3, 4},
			},
			0,
		},
		{
			"edge casee",
			args{
				[]int{1},
			},
			0,
		},
		{
			"equal casee",
			args{
				[]int{1, 2, 3, 3, 3},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.args.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
