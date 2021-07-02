/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  5:33 PM  1/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"edge case",
			args{
				")(",
			},
			[]string{""},
		},
		{
			"left edge case",
			args{
				"x(",
			},
			[]string{"x"},
		},
		{
			"simple test",
			args{
				"()())()",
			},
			[]string{"(())()", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeInvalidParentheses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countInvalid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			"simple case",
			args{
				")(",
			},
			1,
			1,
		},
		{
			"left edge case",
			args{
				"x(",
			},
			1,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countInvalid(tt.args.s)
			if got != tt.want {
				t.Errorf("countInvalid() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countInvalid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
