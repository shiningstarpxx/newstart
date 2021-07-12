/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:21 AM  14/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		l []int
		r []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name : "merge sucess",
			args: args{
				l: []int{1, 3, 5, 7, 9},
				r: []int{2, 4, 6, 8 ,10},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
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
			name : "divide and conque",
			args : args {
				nums: []int{2, 1, 4, 3, 9, 7, 8, 6, 5},
			},
			want : []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}