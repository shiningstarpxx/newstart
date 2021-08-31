/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dfs_lt797_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/25 7:51 PM
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
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
				[][]int{
					{1, 2}, {3}, {3}, {},
				},
			},
			[][]int{{0, 1, 3}, {0, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
