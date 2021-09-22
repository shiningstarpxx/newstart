/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: linklist_lt725_test.go
 * @Version: 1.0.0
 * @Data: 2021/9/22 7:38 AM
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_calcLength(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple test",
			args{
				&ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							3,
							nil,
						},
					},
				},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcLength(tt.args.head); got != tt.want {
				t.Errorf("calcLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitListToParts(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	t1 := &ListNode{
		1,
		nil,
	}
	t2 := &ListNode{
		2,
		nil,
	}
	t3 := &ListNode{
		3,
		nil,
	}
	t1.Next = t2
	t2.Next = t3
	tests := []struct {
		name string
		args args
		want []*ListNode
	}{
		{
			"simple test",
			args{
					t1,
				5,
				},
			[]*ListNode{t1, t2, t3, nil, nil},
			},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitListToParts(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}