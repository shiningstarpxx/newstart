/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_lt600_test.go
 * @Version: 1.0.0
 * @Data: 2021/9/11 8:37 PM
 */
package leetcode

import "testing"

func Test_findIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple test",
			args{
				5,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIntegers(tt.args.n); got != tt.want {
				t.Errorf("findIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
