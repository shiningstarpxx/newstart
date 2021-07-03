/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:16 AM  3/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "testing"

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackCheck(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"simple test",
			args{
				"()",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stackCheck(tt.args.s); got != tt.want {
				t.Errorf("stackCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
