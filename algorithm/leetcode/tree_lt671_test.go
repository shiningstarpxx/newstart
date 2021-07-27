/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: tree_lt671_test.go
 * @Version: 1.0.0
 * @Data: 2021/7/27 1:21 PM
 */
package leetcode

import "testing"

func Test_findSecondMinimumValue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"simple test",  // [2,2,5,null,null,5,7]
			args{
				&TreeNode{
					2,
					&TreeNode{
						2,
						nil,
						nil,
					},
					&TreeNode{
						5,
						&TreeNode{
							5,
							nil,
							nil,
						},
						&TreeNode{
							7,
							nil,
							nil,
						},
					},
				},
			},
			5,
		},
		{
			"bad case",
			args{
				&TreeNode{
					2,
					&TreeNode{
						2,
						nil,
						nil,
					},
					&TreeNode{
						2,
						nil,
						nil,
					},
				},
			},
			-1,
		},
		{
			"bad case 2",  // [2,2,2147483647]
			args{
				&TreeNode{
					2,
					&TreeNode{
						2,
						nil,
						nil,
					},
					&TreeNode{
						2147483647,
						nil,
						nil,
					},
				},
			},
			2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSecondMinimumValue(tt.args.root); got != tt.want {
				t.Errorf("findSecondMinimumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
