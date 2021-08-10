/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: math_lt413_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/10 4:51 PM
 */
package leetcode

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
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
				[]int {1,2,3,4},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
