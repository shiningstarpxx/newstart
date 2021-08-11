/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_lt446_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/11 11:42 PM
 */
package leetcode

import "testing"

func Test_numberOfArithmeticSlices2(t *testing.T) {
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
				[]int{2,4,6,8,10},
			},
			7,
		},
		{
			"all the same",
			args{
				[]int{7,7,7,7,7},
			},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices2(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices2() = %v, want %v", got, tt.want)
			}
		})
	}
}
