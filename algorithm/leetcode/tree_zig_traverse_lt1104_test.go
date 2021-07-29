/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: tree_zig_traverse_lt1104_test.go
 * @Version: 1.0.0
 * @Data: 2021/7/29 2:08 PM
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_pathInZigZagTree(t *testing.T) {
	type args struct {
		label int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"simple case",
			args{
				14,
			},
			[]int{1, 3, 4, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathInZigZagTree(tt.args.label); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathInZigZagTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

//            1
//      3           2
//   4     5     6     7
// 15 14 13 12 11 10  9  8
func Test_reverseLabel(t *testing.T) {
	type args struct {
		row   int
		label int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"second row",
			args{
				2,
				3,
			},
			2,
		},
		{
			"fourth row",
			args{
				4,
				8,
			},
			15,
		},
		{
			"third row",
			args{
				3,
				4,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLabel(tt.args.row, tt.args.label); got != tt.want {
				t.Errorf("reverseLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcRow(t *testing.T) {
	type args struct {
		label int
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
				16,
			},
			5,
		},
		{
			"simple case",
			args{
				14,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcRow(tt.args.label); got != tt.want {
				t.Errorf("calcRow() = %v, want %v", got, tt.want)
			}
		})
	}
}