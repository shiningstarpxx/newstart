/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: binary_search_lt1482_test.go
 * @Version: 1.0.0
 * @Data: 2021/10/7 4:52 PM
 */
package leetcode

import "testing"

func Test_minDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple test",
			args{
				[]int{1,10,3,10,2},
				3,
				1,
			},
			3,
		},
		{
			"another test",
			args{
				[]int{7,7,7,7,12,7,7},
				2,
				3,
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
