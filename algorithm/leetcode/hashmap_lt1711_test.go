/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  2:54 PM  7/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		deliciousness []int
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
				[]int{1,3, 5, 7, 9},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.deliciousness); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
