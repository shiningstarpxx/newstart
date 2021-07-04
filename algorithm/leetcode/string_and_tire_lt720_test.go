/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:27 PM  4/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "testing"

func Test_longestWord(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"simple test",
			args{
				[]string{"w","wo","wor","worl","world"},
			},
			"world",
		},
		{
			"same length",
			args{
				[]string{"a","banana","app","appl","ap","apply","apple"},
			},
			"apple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWord(tt.args.words); got != tt.want {
				t.Errorf("longestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
