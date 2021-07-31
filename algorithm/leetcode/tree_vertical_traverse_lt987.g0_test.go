/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: tree_vertical_traverse_lt987.g0_test.go
 * @Version: 1.0.0
 * @Data: 2021/7/31 5:34 PM
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_verticalTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"simple test",
			args{
				&TreeNode{
					3,
					&TreeNode{
						9,
						nil, nil,
					},
					&TreeNode{
						20,
						&TreeNode{
							15,
							nil,nil,
						},
						&TreeNode{
							7,
							nil, nil,
						},
					},
				},
			},
		[][]int{{9},{3,15},{20},{7}},
		},
		// [1,2,3,4,6,5,7] base case
		// todo(michael): add base
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verticalTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
