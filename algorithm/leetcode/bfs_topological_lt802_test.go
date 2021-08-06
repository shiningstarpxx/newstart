/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: bfs_topological_lt802_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/5 12:38 PM
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_eventualSafeNodes(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"simple test",
			args{
				[][]int{
					{1, 2},
					{2, 3},
					{5},
					{0},
					{5},
					{},
					{},
				},
			},
			[]int{2, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eventualSafeNodes(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventualSafeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
