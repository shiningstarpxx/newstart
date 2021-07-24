/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: brute_force_lt1736_test.go
 * @Version: 1.0.0
 * @Data: 2021/7/24 12:17 PM
 */
package leetcode

import "testing"

func Test_maximumTime(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"simple case",
			args{
				"2?:?0",
			},
			"23:50",
		},
		{
			"second case",
			args{
				"0?:3?",
			},
			"09:39",
		},
		{
			"third case",
			args {
				"1?:22",
			},
			"19:22",
		},
		{
			"bad case",
			args {
				"?4:03",
			},
			"14:03",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTime(tt.args.time); got != tt.want {
				t.Errorf("maximumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
