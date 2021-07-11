/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:38 AM  11/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "testing"

func Test_coinChangeDP(t *testing.T) {
	type args struct {
		coins  []int
		amount int
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
				[]int{1, 2, 5},
				11,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChangeDP(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChangeDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChangeDFS(t *testing.T) {
	type args struct {
		coins  []int
		amount int
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
				[]int{1, 2, 5},
				11,
			},
			3,
		},
		{
			"over time limit",
			args{
				[]int{411,412,413,414,415,416,417,418,419,420,421,422},
				9864,
			},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChangeDFS(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChangeDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}