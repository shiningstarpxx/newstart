/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: sort_squeeze_lt881_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/26 11:21 AM
 */
package leetcode

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple test",
			args{
				[]int{1, 2},
				3,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.args.people, tt.args.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
