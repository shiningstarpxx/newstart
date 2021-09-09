/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: greedy_lt68_test.go
 * @Version: 1.0.0
 * @Data: 2021/9/9 5:58 PM
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"simple test",
			args{
				[]string{"This", "is", "an", "example", "of", "text", "justification."},
				16,
			},
			[]string{"This    is    an", "example  of text", "justification.  "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
