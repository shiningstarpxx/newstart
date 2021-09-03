/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: heap_17.14_test.go
 * @Version: 1.0.0
 * @Data: 2021/9/3 9:48 PM
 */
package interview

import (
	"reflect"
	"testing"
)

func Test_smallestK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"simple test",
			args{
				[]int{1,3,5,7,2,4,6,8},
				4,
			},
			[]int{4,3,2,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestK(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestK() = %v, want %v", got, tt.want)
			}
		})
	}
}
