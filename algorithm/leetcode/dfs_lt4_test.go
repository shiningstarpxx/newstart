/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  12:52 PM  3/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			"simple case",
			args{
				[]int{1, 3},
				[]int{2},
			},
			2,
		},
		{
			"out of index",
			args{
				[]int{1, 2},
				[]int{3, 4},
			},
			2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
