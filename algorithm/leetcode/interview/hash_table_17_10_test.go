/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  1:19 PM  9/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package interview

import "testing"

func Test_majorityElement(t *testing.T) {
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
			"simple case",
			args{
				[]int{1,2,5,9,5,9,5,5,5},
			},
			5,
		},
		{
			"not find",
			args{
				[]int{3,2},
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
