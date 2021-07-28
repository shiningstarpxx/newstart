/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: tree_lt863_test.go
 * @Version: 1.0.0
 * @Data: 2021/7/29 12:03 AM
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_distanceK(t *testing.T) {
	node := &TreeNode{
		5,
		&TreeNode{
			6,
			nil,
			nil,
		},
		&TreeNode{
			2,
			&TreeNode{
				7,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
	}
	type args struct {
		root   *TreeNode
		target *TreeNode
		k      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"simple test",  // 3,5,1,6,2,0,8,null,null,7,4
			args{
				&TreeNode{
					3,
					node,
					&TreeNode{
						1,
						&TreeNode{
							0,
							nil,
							nil,
						},
						&TreeNode{
							8,
							nil,
							nil,
						},
					},
				},
				node,
				2,
			},
			[]int{7, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceK(tt.args.root, tt.args.target, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}
