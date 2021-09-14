/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: string_lt524_test.go
 * @Version: 1.0.0
 * @Data: 2021/9/14 11:08 PM
 */
package leetcode

import "testing"

func Test_findLongestWord(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"simple test",
			args{
				"abpcplea",
				[]string{"ale","apple","monkey","plea"},
			},
			"apple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestWord(tt.args.s, tt.args.dictionary); got != tt.want {
				t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
