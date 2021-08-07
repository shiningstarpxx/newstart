/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: two_pointers_slow_fast_lt457_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/7 11:55 AM
 */
package leetcode

import "testing"

func Test_circularArrayLoop(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"simple test",
			args{
				[]int{2,-1,1,2,2},
			},
			true,
		},
		{
			"bad case",
			args{
				[]int{-2,1,-1,-2,-2},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayLoop(tt.args.nums); got != tt.want {
				t.Errorf("circularArrayLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
