/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: hash_map_lt1436_test.go
 * @Version: 1.0.0
 * @Data: 2021/10/1 10:12 AM
 */
package leetcode

import "testing"

func Test_destCityHelper(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"simple test",
			args{
				[][]string{
					{"B", "C"}, {"D", "B"}, {"C", "A"},
				},
			},
			"A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destCityHelper(tt.args.paths); got != tt.want {
				t.Errorf("destCityHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}
