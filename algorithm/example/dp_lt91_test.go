/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  6:36 PM  27/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package example

import "testing"

func Test_calcOne91(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcOne91(tt.args.c); got != tt.want {
				t.Errorf("calcOne91() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcTwo91(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcTwo91(tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("calcTwo91() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDecodings91(t *testing.T) {
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
				s : "12",
			},
			want : 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings91(tt.args.s); got != tt.want {
				t.Errorf("numDecodings91() = %v, want %v", got, tt.want)
			}
		})
	}
}
