/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  4:58 PM  16/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package interview

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				[]int {5,7,7,8,8,10},
				8,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
