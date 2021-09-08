/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: heap_greedy_lt502_test.go
 * @Version: 1.0.0
 * @Data: 2021/9/8 10:27 PM
 */
package leetcode

import "testing"

func Test_findMaximizedCapital(t *testing.T) {
	type args struct {
		k       int
		w       int
		profits []int
		capital []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple test",
			args{
				2,
				0,
				[]int{1, 2, 3},
				[]int{0, 1, 1},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximizedCapital(tt.args.k, tt.args.w, tt.args.profits, tt.args.capital); got != tt.want {
				t.Errorf("findMaximizedCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
