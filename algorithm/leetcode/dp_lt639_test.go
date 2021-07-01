/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:54 AM  27/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "testing"

func Test_calcOne(t *testing.T) {
	type args struct {
		c byte
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
				c : '*',
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcOne(tt.args.c); got != tt.want {
				t.Errorf("calcOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcTwo(t *testing.T) {
	type args struct {
		b byte
		c byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name : "1*",
			args : args{
				b : '1',
				c : '*',
			},
			want: 9,
		},
		{
			name : "2*",
			args : args{
				b : '2',
				c : '*',
			},
			want: 6,
		},
		{
			name : "**",
			args : args{
				b : '*',
				c : '*',
			},
			want: 15,
		},
		{
			name : "0*, 00",
			args : args{
				b : '0',
				c : '*',
			},
			want: 0,
		},
		{
			name : "**",
			args : args{
				b : '*',
				c : '1',
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcTwo(tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("calcTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
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
				s : "*",
			},
			want: 9,
		},
		{
			name : "edge case",
			args : args{
				s : "*1",
			},
			want: 11,
		},
		{
			name : "edge case 2",
			args : args{
				s : "2839",
			},
			want: 1,
		},
		{
			name : "need mod",
			args : args{
				s : "*********",
			},
			want: 291868912,  // if not mod will be 1291868919
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
