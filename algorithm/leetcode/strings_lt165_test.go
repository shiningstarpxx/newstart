/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: strings_lt165_test.go
 * @Version: 1.0.0
 * @Data: 2021/9/1 10:36 AM
 */
package leetcode

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple test",
			args{
				"1.01",
				"1.001",
			},
			0,
		},
		{
			"augly test",
			args{
				"0.1",
				"1.1",
			},
			-1,
		},
		{
			"edge case",
			args{
				"1.0.1",
				"1",
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
