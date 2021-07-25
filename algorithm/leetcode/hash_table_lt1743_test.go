/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: hash_table_lt1743_test.go
 * @Version: 1.0.0
 * @Data: 2021/7/25 4:18 PM
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_restoreArray(t *testing.T) {
	type args struct {
		adjacentPairs [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases. [2,1],[3,4],[3,2]
		{
			"simple case",
			args{
				[][]int{
					{2, 1},
					{3, 4},
					{3,2},
				},
			},
			[]int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreArray(tt.args.adjacentPairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
