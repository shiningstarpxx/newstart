/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:37 AM  6/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

import "testing"

func TestLowerBound(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "LowerBound",
			args: args{
				arr: []int{5,5,5,6,6,6,7,7},
				target: 6,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerBound(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("LowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUppperBound(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "UpperBound",
			args: args{
				arr: []int{5,5,5,6,6,6,7,7},
				target: 6,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UppperBound(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("UppperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceLowerBound(t *testing.T) {
	type args struct {
		arr    []int
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
				[]int{5, 27, 23, 234,32,423, 23234,23,2,34,23,23,4,234},
				6,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceLowerBound(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("sliceLowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}