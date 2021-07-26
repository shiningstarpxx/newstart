/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_lt300_test.go
 * @Version: 1.0.0
 * @Data: 2021/7/26 12:35 PM
 */
package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
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
				[]int{10,9,2,5,3,7,101,18},
			},
			4,
		},
		{
			"bad case",
			args{
				[]int{1,3,6,7,9,4,10,5,6},
			},
			6,
		},
		{
			"all the same",
			args{
				[]int{7,7,7,7,7,7,7},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
