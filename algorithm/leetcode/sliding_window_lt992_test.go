/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:55 PM  12/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "testing"

func Test_subarraysWithKDistinct(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				[]int{1,2,1,2,3},
				2,
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysWithKDistinct(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
