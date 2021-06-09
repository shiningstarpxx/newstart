/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  1:08 PM  9/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewSTree(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *STree
	}{
		// TODO: Add test cases.
		{
			name: "Segment Tree",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 9, 10},
			},
			want: nil,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSTree(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				a := assert.New(t)
				v, e := got.Query(1, 3)
				a.Nil(e)
				a.Equal(5, v)

				got.UpdateTreeNode(2, 1)
				v, e = got.Query(1, 3)
				a.Nil(e)
				a.Equal(3, v)
			} else {
				t.Errorf("NewSTree() = %v, want %v", got, tt.want)
			}
		})
	}
}