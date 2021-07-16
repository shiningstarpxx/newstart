/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  3:28 PM  15/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "testing"

func Test_maximumElementAfterDecrementingAndRearranging(t *testing.T) {
	type args struct {
		arr []int
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
				[]int{2,2,1,2,1},
			},
			2,
		},
		{
			"bad case, need first must one",
			args{
				[]int{100,1,1000},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumElementAfterDecrementingAndRearranging(tt.args.arr); got != tt.want {
				t.Errorf("maximumElementAfterDecrementingAndRearranging() = %v, want %v", got, tt.want)
			}
		})
	}
}
