/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_LIS_lt673_test.go
 * @Version: 1.0.0
 * @Data: 2021/9/20 5:45 PM
 */
package algorithm

import (
	"strings"
	"testing"
)

func Test_findNumberOfLIS(t *testing.T) {
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
				[]int{1, 3, 5, 4, 7},
			},
			2,
		},
		{
			"special case",
			args{
				[]int{2, 2, 2, 2, 2},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
