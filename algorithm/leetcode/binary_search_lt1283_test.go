/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: binary_search_lt1283_test.go
 * @Version: 1.0.0
 * @Data: 2021/10/13 9:25 PM
 */
package leetcode

import "testing"

func Test_smallestDivisor(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple test",
			args{
				[]int{1, 2, 5, 9},
				6,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestDivisor(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("smallestDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}
