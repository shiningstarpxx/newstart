/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  3:20 PM  26/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"testing"
)

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name : "simple test",
			args : args{
				beginWord: "hit",
				endWord: "cog",
				wordList: []string{"hot","dot","dog","lot","log","cog"},
			},
			want : 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
