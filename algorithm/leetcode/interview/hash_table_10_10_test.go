/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  1:44 PM  18/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package interview

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			"simple test",
			args{
				[]string {"eat","tea","tan","ate","nat","bat"},
			},
			[][]string{
				{"ate","eat","tea"},
				{"nat","tan"},
				{"bat"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
