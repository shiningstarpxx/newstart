/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  3:55 PM  4/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"simple test",
			args{
				[]int{1, 2, 2, 4},
			},
			[]int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorNums(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitCalc(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"simple test",
			args{
				[]int{1, 2, 2, 4},
			},
			[]int{2, 3},
		},
		{
			"edge case",
			args{
				[]int{3,2,3,4,6,5},
			},
			[]int{3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitCalc(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bitCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}