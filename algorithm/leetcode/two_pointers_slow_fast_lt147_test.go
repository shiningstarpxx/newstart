/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: two_pointers_slow_fast_lt147_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/7 7:21 PM
 */
package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	var a *ListNode
	a = &ListNode{
		2,
		&ListNode{
			0,
			&ListNode{
				4,
				a,
			},
		},
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"simple test",
			args{
				nil,
			},
			false,
		},
		{  // todo(michael): 这个test错误
			"circle test",
			args{
				&ListNode{
						3,
						a,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
